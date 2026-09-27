# Wristrunner OS - Installation Guide

This guide explains on how to use the OS on a QEMU virtual machine.

# First Time Start Up

Before starting, remember to clone the github repo and to download [qemu-rootfs.qcow2](https://github.com/MrDaPoyo/Wristrunner/releases/download/v1.0.0/qemu-rootfs.qcow2.tar.gz) from Github Releases.

'''bash
git clone [https://github.com/MrDaPoyo/Wristrunner.git](https://github.com/MrDaPoyo/Wristrunner.git)
cd Wristrunner/OS

After entering the OS directory, move the [qemu-rootfs.qcow2](https://github.com/MrDaPoyo/Wristrunner/releases/download/v1.0.0/qemu-rootfs.qcow2.tar.gz) file here.
## Windows with WSL

Open Powershell as administrator and run:

'''bash
wsl --install

Restart PC and then open Ubuntu from Windows Start Menu
Install QEMU dependencies:

'''bash
sudo apt update && sudo apt install -y qemu-system-arm
cd ~/Wristrunner/OS
./run.sh

## Windows without WSL

'''bash
winget install QEMU.QEMU
cd ~/Wristrunner/OS
./run.bat

## Ubuntu / Debian

'''bash
sudo apt update && sudo apt install -y qemu-system-arm
cd ~/Wristrunner/OS
./run.sh

## Arch Linux

'''bash
sudo pacman -S qemu-emulators-full
cd ~/Wristrunner/OS
./run.sh
