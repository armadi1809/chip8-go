package main

import (
	"log"
	"os"

	"github.com/armadi1809/chip8-go/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Chip 8 In Golang")
	var romPath string
	if len(os.Args) > 2 {
		panic("invalid arguments, use a rom file path as the only argument to the program or run without arguments to use the emulator built in games")
	}

	if len(os.Args) == 2 {
		romPath = os.Args[1]
	}

	game := graphics.NewGame(romPath)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
