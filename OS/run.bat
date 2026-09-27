@echo off
setlocal

where qemu-system-aarch64.exe >nul 2>nul
if %errorlevel% equ 0 (
    set QEMU_CMD=qemu-system-aarch64.exe
) else if exist "C:\Program Files\qemu\qemu-system-aarch64.exe" (
    set QEMU_CMD="C:\Program Files\qemu\qemu-system-aarch64.exe"
) else (
    echo [ERROR] QEMU ARM64 was not found!
    pause
    exit /b 1
)

IF NOT EXIST "qemu-rootfs.qcow2" (
    echo [ERROR] qemu-rootfs.qcow2 not found!
    pause
    exit /b 1
)

echo Launching Wristrunner OS on QEMU ARM64...
%QEMU_CMD% ^
  -M virt -cpu cortex-a72 -m 2G -smp 4 ^
  -kernel precompiled\Image ^
  -append "console=ttyAMA0 console=tty1 root=/dev/vda1 rw rootwait rootfstype=ext4" ^
  -drive file=qemu-rootfs.qcow2,format=qcow2,if=virtio ^
  -netdev user,id=net0 -device virtio-net-pci,netdev=net0 ^
  -device virtio-gpu-pci ^
  -device virtio-keyboard-pci ^
  -device virtio-tablet-pci ^
  -serial stdio

pause