# libmakepkg drop-ins

This repo contains files that can be dropped into libmakepkg to extend the functionality of makepkg.

It is intended that individual files from this repo are installed to meet a user's or distribution's needs.

## Contents

### autodep/

Automatically add dependencies and provides to packages

- *binary_provides.sh*: add a package's executables into its provisions.  The searched directories are controlled by the `BIN_DIRS` configuration variable.  The entries in the `BIN_DIRS` array need the format `prefix:path`. For example, `BIN_DIRS=(cmd:usr/bin)` would add `/usr/bin/sh` to the provides entry as `cmd:sh`. (compatibility: pacman>=6.1.0)

### reproducible/

Setting environment variables that improve package reproducibility

- *python.sh*: disable hash randomization when creating .pyc files (compatibility: pacman>=7.1.0)

### tidy/

Use to modify files after the package() function has been run.

- *50-pestrip.sh*: Strip debug symbols from Portable Executable (PE) format files.  Enabled with the `pestrip` option. (compatibility: pacman>=7.1.0)

- *10-ruby.sh*: Remove unreproducible (and unneeded) files created during ruby packaging. (compatibility: pacman>=5.0.0)

- *50-zipkmod.sh*: Compress kernel modules via zstd. Enabled with the ``zipkmod`` option. (compatibility: pacman>=5.0.0)
