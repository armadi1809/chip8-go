package main

import (
	"log"
	"os"

	"github.com/armadi1809/chip8-go/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid Arguments. Run the emulator with one rom file path")
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Chip 8 In Golang")
	game := graphics.NewGame(os.Args[1])
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
