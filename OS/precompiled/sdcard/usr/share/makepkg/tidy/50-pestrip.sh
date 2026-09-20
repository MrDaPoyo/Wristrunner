#!/bin/bash
#
#   pestrip.sh - Strip debugging symbols from PE format files
#
#   Copyright (c) 2024 Pacman Development Team <pacman-dev@lists.archlinux.org>
#
#   This program is free software; you can redistribute it and/or modify
#   it under the terms of the GNU General Public License as published by
#   the Free Software Foundation; either version 2 of the License, or
#   (at your option) any later version.
#
#   This program is distributed in the hope that it will be useful,
#   but WITHOUT ANY WARRANTY; without even the implied warranty of
#   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#   GNU General Public License for more details.
#
#   You should have received a copy of the GNU General Public License
#   along with this program.  If not, see <http://www.gnu.org/licenses/>.
#

[[ -n "$LIBMAKEPKG_TIDY_PESTRIP_SH" ]] && return
LIBMAKEPKG_TIDY_PESTRIP_SH=1

MAKEPKG_LIBRARY=${MAKEPKG_LIBRARY:-'/usr/share/makepkg'}

source "$MAKEPKG_LIBRARY/util/message.sh"
source "$MAKEPKG_LIBRARY/util/option.sh"
source "$MAKEPKG_LIBRARY/tidy/50-strip.sh"

packaging_options+=('pestrip')
tidy_modify+=('tidy_strip_pe')


build_id_pe() {
	local bid=`objdump -p "$1" | sed -n "/RSDS signature / {s/.*signature //; s/ age.*//p; q;}"`
	echo ${bid:6:2}${bid:4:2}${bid:2:2}${bid:0:2}${bid:10:2}${bid:8:2}${bid:14:2}${bid:12:2}${bid:16}
}

# placeholder - currently not supported as solutions
# like llvm-dwarfdump bring heavy dependencies.
package_source_files_pe() {
	:
}

collect_debug_symbols_pe() {
	local binary=$1; shift

	if check_option "debug" "y"; then
		local bid=$(build_id_pe "$binary")

		# has this file already been stripped
		if [[ -n "$bid" ]]; then
			if [[ -f "$dbgdir/.build-id/${bid:0:2}/${bid:2}.debug" ]]; then
				return
			fi
		elif [[ -f "$dbgdir/$binary.debug" ]]; then
			return
		fi

		# copy source files to debug directory
		package_source_files_pe "$binary"

		# copy debug symbols to debug directory
		mkdir -p "$dbgdir/${binary%/*}"

		# abandon processing files that are not a recognised format
		if ! objcopy --only-keep-debug "$binary" "$dbgdir/$binary.debug" 2>/dev/null; then
			return
		fi

		safe_objcopy "$binary" --remove-section=.gnu_debuglink
		safe_objcopy "$binary" --add-gnu-debuglink="$dbgdir/${binary#/}.debug"

		if [[ -n "$bid" ]]; then
			local target
			mkdir -p "$dbgdir/.build-id/${bid:0:2}"

			target="../../../../../${binary#./}"
			target="${target/..\/..\/usr\/lib\/}"
			target="${target/..\/usr\/}"
			ln -s "$target" "$dbgdir/.build-id/${bid:0:2}/${bid:2}"

			target="../../${binary#./}.debug"
			ln -s "$target" "$dbgdir/.build-id/${bid:0:2}/${bid:2}.debug"
		fi
	fi
}


process_file_stripping_pe() {
	local binary="$1"
	local file_type

	file_type=$(LC_ALL=C file --no-sandbox "$binary" 2>/dev/null)

	if [[ "$file_type" =~ .*PE32[+]?\ executable.* ]]; then
		collect_debug_symbols_pe "$binary"
		safe_strip_file "$binary" $STRIP_SHARED
	fi
}

tidy_strip_pe() {
	if check_option "pestrip" "y"; then
		msg2 "$(gettext "Stripping unneeded symbols from PE format files...")"
		# make sure library stripping variables are defined to prevent excess stripping
		[[ -z ${STRIP_SHARED+x} ]] && STRIP_SHARED="-S"

		if check_option "debug" "y"; then
			dbgdir="$pkgdirbase/$pkgbase-debug/usr/lib/debug"
			dbgsrcdir="${DBGSRCDIR:-/usr/src/debug}/${pkgbase}"
			dbgsrc="$pkgdirbase/$pkgbase-debug$dbgsrcdir"
			mkdir -p "$dbgdir" "$dbgsrc"
		fi

		_parallel_stripper() {
			# Inherit traps in subshell to perform cleanup after an interrupt
			set -E
			(
				local jobs binary

				while IFS= read -rd '' binary ; do
					# Be sure to keep the number of concurrently running processes less
					# than limit value to prevent an accidental fork bomb.
					jobs=($(jobs -p))
					(( ${#jobs[@]} >= $NPROC )) && wait -n "${jobs[@]}"

					process_file_stripping_pe "$binary" &
				done < <(find . -type f -perm -u+w -links 1 -print0 2>/dev/null)

				# Wait for all jobs to complete
				wait
			)
			set +E
		}
		_parallel_stripper

		# hardlinks only need processed once, but need additional links in debug packages
		declare -A hardlinks
		while IFS= read -rd '' binary ; do
			if check_option "debug" "y"; then
				local inode="$(stat -c '%i %n' -- "$binary")"
				inode=${inode%% *}
				if [[ -z "${hardlinks[$inode]}" ]]; then
					hardlinks[$inode]="$binary"
				else
					if [[ -f "$dbgdir/${hardlinks[$inode]}.debug" ]]; then
						mkdir -p "$dbgdir/${binary%/*}"
						ln "$dbgdir/${hardlinks[$inode]}.debug" "$dbgdir/${binary}.debug"
						continue
					fi
				fi
			fi
			process_file_stripping_pe "$binary"
		done < <(find . -type f -perm -u+w -links +1 -print0 2>/dev/null | LC_ALL=C sort -z)
	fi
}
