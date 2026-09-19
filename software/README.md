# Wristrunner Launcher
Uses [Gio](https://gioui.org/doc/install/linux)

## Installation
Gio requires 
```sh
sudo pacman -S --needed vulkan-headers vulkan-tools
go install gioui.org/cmd/gogio@latest
```

## App development
Every app needs a `manifest.wr` file. That file should contain the following info:

- `author`: Who made the app?
- `name`: What's the app's name?
- `version`: Self-explanatory.

An example `manifest.wr` file looks like this:

```toml
author: James Bond
name: Test
version: 1.0.1
```

## Layout
```
software/
├── go.work (make everything a single workspace)
├── uikit/              
│   ├── go.mod
│   ├── theme.go
│   ├── statusbar.go (e.g)
│   └── list.go (e.g)
├── launcher/
│   ├── go.mod
│   ├── main.go (entrypoint)
│   └── apps.go (e.g)
└── apps/
    ├── notes/
    │   ├── go.mod (e.g)
    │   └── main.go (e.g)
    └── scanner/
        ├── go.mod (e.g)
        └── main.go (e.g)
```