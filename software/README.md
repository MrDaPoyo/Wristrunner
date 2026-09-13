# Wristrunner Launcher
Uses [Gio](https://gioui.org/doc/install/linux)

## Installation
Gio requires 
```sh
sudo pacman -S --needed vulkan-headers vulkan-tools
go install gioui.org/cmd/gogio@latest
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