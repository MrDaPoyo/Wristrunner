#!/usr/bin/bash
#
#   ruby.sh - Remove unreproducible files from ruby packages
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

[[ -n "$LIBMAKEPKG_TIDY_RUBY_SH" ]] && return
LIBMAKEPKG_TIDY_RUBY_SH=1

MAKEPKG_LIBRARY=${MAKEPKG_LIBRARY:-'/usr/share/makepkg'}

source "$MAKEPKG_LIBRARY/util/message.sh"
source "$MAKEPKG_LIBRARY/util/option.sh"


tidy_remove+=('tidy_ruby')

tidy_ruby() {
	if ! type -p gem >/dev/null; then
		return
	fi

	msg2 "$(gettext "Purging unreproducible ruby files...")"

	local gd="$pkgdir/$(gem env gemdir)"

	rm -fr \
		"${gd}"/cache/ \
		"${gd}"/gems/*/vendor/ \
		"${gd}"/doc/*/ri/ext/


	find "${gd}/gems/" \
		-type f \
			\( \
			-iname "*.o" -o \
			-iname "*.c" -o \
			-iname "*.so" -o \
			-iname "*.time" -o \
			-iname "gem.build_complete" -o \
			-iname "Makefile" \
		\) \
		-delete 2>/dev/null

	find "${gd}/extensions/" \
		-type f \
		\( \
			-iname "mkmf.log" -o \
			-iname "gem_make.out" \
		\) \
		-delete 2>/dev/null

}
