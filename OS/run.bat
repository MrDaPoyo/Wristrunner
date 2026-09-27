@echo off
IF NOT EXIST "qemu-rootfs.qcow2" (
   echo qemu-rootfs.qcow2 not found!
   echo Please download qemu-rootfs.qcow2 from Github Releases and place it in this directory.
   pasue
   exit /b 1
)

echo Launching Wristrunner OS on QEMU ARM64...
qemu-system-aarch64.exe ^
 -M virt -cpu cortex-a72 -m 2G -smp 4 ^
 -kernel precompiled\Image ^
 -append "console=ttyAMA0 console=tty1 root=/dev/vda1 rw rootwait" ^
 -drive if=none,file=qemu-rootfs.qcow2,format=qcow2,id=hd0 ^
 -netdev user,id=net0 -device virtio-net-device,netdev=net0 ^
 -device virtio-gpu-pci ^
 -device virtio-keyboard-pci ^
 -device virtio-tablet-pci ^
 -serial stdio

pause
