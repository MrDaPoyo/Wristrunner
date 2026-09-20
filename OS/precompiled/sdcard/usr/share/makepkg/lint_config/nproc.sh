#!/usr/bin/bash
#
#   nproc.sh - Check that NPROC setting is valid
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

[[ -n $LIBMAKEPKG_LINT_CONFIG_NPROC_SH ]] && return
LIBMAKEPKG_LINT_CONFIG_NPROC_SH=1

MAKEPKG_LIBRARY=${MAKEPKG_LIBRARY:-'/usr/share/makepkg'}

source "$MAKEPKG_LIBRARY/util/message.sh"

lint_config_functions+=('lint_nproc')


lint_nproc() {
	# provide default value if NPROC is not set
	[[ -z $NPROC ]] && NPROC=1

	if [[ $NPROC = *[^[:digit:]]* ]]; then
		error "$(gettext "%s contains invalid characters: %s")" \
			"\$NPROC" "${NPROC//[[:digit:]]}"
		return 1
	fi
}
