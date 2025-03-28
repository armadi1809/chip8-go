package main

import (
	"log"

	"github.com/armadi1809/chip8-go/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Chip 8 In Golang")
	game := graphics.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
