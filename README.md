# Chip8-Go: A CHIP-8 Emulator in Go

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A CHIP-8 virtual machine emulator written in the Go programming language, using the [Ebitengine](https://ebitengine.org/) 2D game library for graphics, input, and sound.

This project is a reasonably accurate implementation of the CHIP-8 interpreter, suitable for playing classic CHIP-8 games and homebrew ROMs. At its current state, the code is passing most of the tests in the popular [Timendus' test suite for the CHIP 8](https://github.com/Timendus/chip8-test-suite). It also served as personal learning exercise in emulation, Go, and Ebitengine.

## Features

- CHIP-8 CPU Emulation
- 64x32 Monochrome Graphics Display (via Ebitengine)
- Keyboard Input Mapping (via Ebitengine)
- Basic Sound Output (BEEP instruction via Ebitengine)

## Prerequisites

- **Go:** Version 1.24 or later recommended. ([Installation Guide](https://go.dev/doc/install))
- **Git:** For cloning the repository.
- **Platform-Specific Dependencies for Ebitengine:** Ebitengine requires certain libraries depending on your operating system (e.g., C compilers, graphics/audio libraries). Please refer to the [Ebitengine Installation Guide](https://ebitengine.org/en/documents/install.html) for details specific to your OS (Linux, macOS, Windows).

## Installation & Building

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/armadi1809/chip8-go.git
    cd chip8-go
    ```

2.  **Build the executable:**
    ```bash
    go build -o chip8emulator ./cmd
    ```
    This will create an executable file named `chip8emulator` (or `chip8emulator.exe` on Windows) in the current directory.

## Usage

Run the emulator from your terminal, providing the path to a CHIP-8 ROM file as an argument:

```bash
./chip8emulator path/to/your/romfile.ch8
```

There are a couple of ROMS available as part of this repo under /roms.

Alternatively you can run the program without any roms and you can use a couple of embedded roms selectable via a dropdown.

```bash
./chip8emulator
```

### Online Demo

You can also try the emulator directly in your web browser by visiting:
[https://azizrmadi.com/projects/chip8/](https://azizrmadi.com/projects/chip8/)

The web version includes some ROMs selectable from a dropdown and features as the desktop version.
