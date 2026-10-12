#!/usr/bin/env bash
set -e

# Download rootfs if not downloaded
if [ ! -f "qemu-rootfs.qcow2" ]; then
    echo "qemu-rootfs.qcow2 not found!"
    echo "Download qemu-rootfs from Github Releases and place it in this folder"
    exit 1

fi

echo "Launching Wristrunner OS on QEMU ARM64"
qemu-system-aarch64 \
  -M virt -cpu cortex-a72 -m 2G -smp 4 \
  -kernel precompiled/Image \
  -append "console=ttyAMA0 console=tty1 quiet loglevel=3 systemd.show_status=0 root=/dev/vda1 rw rootwait rootfstype=ext4" \
  -drive file=qemu-rootfs.qcow2,format=qcow2,if=virtio \
  -netdev user,id=net0 -device virtio-net-pci,netdev=net0 \
  -device virtio-gpu-pci \
  -device virtio-keyboard-pci \
  -device virtio-tablet-pci \
  -serial stdio
