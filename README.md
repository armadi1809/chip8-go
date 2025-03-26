# Chip8-Go: A CHIP-8 Emulator in Go

<!-- Optional: Add a cool banner/logo here -->

<!-- [![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/gochip-8)](https://goreportcard.com/report/github.com/yourusername/gochip-8) Replace with your actual repo path -->
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A CHIP-8 virtual machine emulator written in the Go programming language, using the [Ebitengine](https://ebitengine.org/) 2D game library for graphics, input, and sound.

This project is a reasonably accurate implementation of the CHIP-8 interpreter, suitable for playing classic CHIP-8 games and homebrew ROMs. At its current state, the code is passing most of the tests in the popular [Timendus' test suite for the CHIP 8](https://github.com/Timendus/chip8-test-suite). It also served as personal learning exercise in emulation, Go, and Ebitengine.

<!-- Recommended: Add a screenshot or GIF of the emulator running a game -->
<!-- ![Screenshot of GoCHIP-8 running Pong](docs/screenshot.png) -->

## Features

*   CHIP-8 CPU Emulation
*   64x32 Monochrome Graphics Display (via Ebitengine)
*   Keyboard Input Mapping (via Ebitengine)
*   [Work In Progress] Basic Sound Output (BEEP instruction via Ebitengine)
*   Configurable Clock Speed
*   Adjustable Display Scaling

## Prerequisites

*   **Go:** Version 1.18 or later recommended. ([Installation Guide](https://go.dev/doc/install))
*   **Git:** For cloning the repository.
*   **Platform-Specific Dependencies for Ebitengine:** Ebitengine requires certain libraries depending on your operating system (e.g., C compilers, graphics/audio libraries). Please refer to the [Ebitengine Installation Guide](https://ebitengine.org/en/documents/install.html) for details specific to your OS (Linux, macOS, Windows).

## Installation & Building

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/armadi1809/chip8-go.git # Replace with your repo URL
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