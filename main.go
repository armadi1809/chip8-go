package main

import (
	"fmt"

	"github.com/armadi1809/chip8-go/chip8"
)

func main() {
	emulator := chip8.New()
	emulator.Initialize()

	fmt.Println("Running the Chip 8 emulator")
}
