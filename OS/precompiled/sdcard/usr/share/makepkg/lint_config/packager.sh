#!/usr/bin/bash
#
#   source_date_epoch.sh - Check that reproducible builds timestamp is valid
#
#   Copyright (c) 2018-2025 Pacman Development Team <pacman-dev@lists.archlinux.org>
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

[[ -n $LIBMAKEPKG_LINT_CONFIG_PACKAGER_SH ]] && return
LIBMAKEPKG_LINT_CONFIG_PACKAGER_SH=1

MAKEPKG_LIBRARY=${MAKEPKG_LIBRARY:-'/usr/share/makepkg'}

source "$MAKEPKG_LIBRARY/util/message.sh"

lint_config_functions+=('lint_packager')


lint_packager() {
	# pacman relies on this format to extract an email address from PACKAGER for WKD key lookup
	local match='^([^<>]+ )?<[^<>]*>$'
	if ! [[ $PACKAGER == "Unknown Packager" || $PACKAGER =~ $match ]]; then
		warning "$(gettext "PACKAGER should have the format 'Example Name <email@address.invalid>'")"
	fi
}
