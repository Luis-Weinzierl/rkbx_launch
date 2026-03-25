# Rkbx Launch
Rkbx Launch is a Launcher for [rkbx_link](https://github.com/grufkork/rkbx_link) written in Go and fyne and thus raises minimal overhead. It simplifies configuration by abstracting the old config-file behind a specially-built UI and makes working with rkbx_link just that bit more pleasant.

## Developer Guide
### Prerequisites
To build and test rkbx_launch, you'll need to have Go, GCC and the Fyne Toolchain installed:
- Go: Please follow the official guide [here](https://go.dev/doc/install).
- GCC: This is best done using MSYS2 MinGW:
    - Get [MSYS2](https://www.msys2.org/)
    - Open MinGW 64-bit 
    - Run these commands to install gcc:
    ```sh
    pacman -Syu
    pacman -S git mingw-w64-x86_64-toolchain mingw-w64-x86_64-go
    ```
    - Add `C:\msys64\mingw64\bin` to your PATH
- Fyne: You'll need the fyne helper to build rkbx_link:
    - Run the following command to install the fyne helper:
    ```
    go install fyne.io/tools/cmd/fyne@latest
    ```
    - If not already present, add `%USERPROFILE%\Go\bin` to your PATH
    - Check your installation by using the official [Fyne Setup](https://geoffrey-artefacts.fynelabs.com/github/andydotxyz/fyne-io/setup/latest/) (PS.: doesn't set up anything but rather checks if everything is set up)

## Running
To run rkbx_launch (e.g. for testing), run the following command in the project root:
```sh
go run .
```

## Building
To build and package rkbx_launch, run this command:
```sh
fyne package
```

NOTE: rkbx_link requires the `/assets` and `/rkbx_link` folders at runtime.