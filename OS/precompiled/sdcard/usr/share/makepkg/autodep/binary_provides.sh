#!/bin/bash
#
#   binary_provides.sh - Automatically add a package's binaries to provides
#
#   Copyright (c) 2025 Pacman Development Team <pacman-dev@lists.archlinux.org>
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

[[ -n "$LIBMAKEPKG_AUTODEP_BINARY_PROVIDES_SH" ]] && return
LIBMAKEPKG_AUTODEP_BINARY_PROVIDES_SH=1

MAKEPKG_LIBRARY=${MAKEPKG_LIBRARY:-'/usr/share/makepkg'}

autodep_functions+=('binary_provides')

binary_provides() {
	if check_option "autodeps" "y"; then
		for bin in ${BIN_DIRS[@]}; do
			dir=${bin#*:}
			prefix=${bin%%:*}

			if [[ ! -d "$pkgdir/$dir" ]]; then
				continue
			fi

			mapfile -t filenames < <(find "$pkgdir/$dir" -maxdepth 1 -type f -follow | LC_ALL=C sort)

			for fn in "${filenames[@]}"; do
				if [[ -x $fn ]]; then
					provides+=("$prefix:${fn##*/}")
				fi
			done
		done
	fi
}
