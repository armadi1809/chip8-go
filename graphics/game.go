package graphics

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"os"

	"github.com/armadi1809/chip8-go/chip8"
	"github.com/armadi1809/chip8-go/graphics/widgets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	emulator            *chip8.Chip8
	pixels              []byte
	beepSoundEffectData []byte
	audioContext        *audio.Context
	combobox            *widgets.ComboBox
}

//go:embed audio/beep-03.mp3
var defaultBeep []byte

// ROMS
//
//go:embed roms/pong2.ch8
var pong2 []byte

//go:embed roms/spaceinvaders.ch8
var spaceinvaders []byte

var titles = []string{"pong2", "spaceinvaders"}

var titleToRomMap = map[string][]byte{
	"pong2":         pong2,
	"spaceinvaders": spaceinvaders,
}

const clockRate = 10

var keyboardToEmulatorMap map[ebiten.Key]int = map[ebiten.Key]int{
	ebiten.Key1: 0x1,
	ebiten.Key2: 0x2,
	ebiten.Key3: 0x3,
	ebiten.Key4: 0xC,

	ebiten.KeyQ: 0x4,
	ebiten.KeyW: 0x5,
	ebiten.KeyE: 0x6,
	ebiten.KeyR: 0xD,

	ebiten.KeyA: 0x7,
	ebiten.KeyS: 0x8,
	ebiten.KeyD: 0x9,
	ebiten.KeyF: 0xE,

	ebiten.KeyZ: 0xA,
	ebiten.KeyX: 0x0,
	ebiten.KeyC: 0xB,
	ebiten.KeyV: 0xF,
}

func (g *Game) Update() error {
	updateKeys(g.emulator)
	for range clockRate {
		g.emulator.EmulateCycle()
	}
	g.emulator.UpdateTimers()
	g.playBeepSoundEffectIfNeeded()
	if g.combobox != nil {
		previousrom := g.combobox.SelectedIndex
		g.combobox.Update()
		if g.combobox.SelectedIndex != previousrom {
			g.emulator.Initialize()
			g.emulator.LoadProgram(titleToRomMap[titles[g.combobox.SelectedIndex]])
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.emulator.DrawFlag {
		g.pixels = getPixelsFromEmulator(g.emulator)
		g.emulator.DrawFlag = false
	}
	emulatorImg := ebiten.NewImage(64, 32)

	emulatorImg.WritePixels(g.pixels)
	emulatorDrawOptions := &ebiten.DrawImageOptions{}
	emulatorDrawOptions.GeoM.Scale(10, 12)
	emulatorDrawOptions.GeoM.Translate(0, 50)
	screen.DrawImage(emulatorImg, emulatorDrawOptions)
	if g.combobox != nil {
		g.combobox.Draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame(romPath string) *Game {
	emulator := chip8.New()
	emulator.Initialize()
	if romPath != "" {
		rom, err := os.ReadFile(romPath)
		if err != nil {
			log.Fatalf("unable to read the rom file path %v", err)
		}
		emulator.LoadProgram(rom)
	} else {
		emulator.LoadProgram(pong2)
	}
	game := &Game{
		emulator:     emulator,
		audioContext: audio.NewContext(48000),
	}
	if romPath != "" {
		game.combobox = nil
	} else {
		game.combobox = widgets.NewComboBox(3, 3, 150, 20, titles)
	}
	game.setBeepSoundEffect()
	return game
}

func updateKeys(emulator *chip8.Chip8) {
	for key, val := range keyboardToEmulatorMap {
		if ebiten.IsKeyPressed(key) || inpututil.IsKeyJustPressed(key) {
			emulator.Key[val] = 1
		} else {
			emulator.Key[val] = 0
		}
	}
}

func getPixelsFromEmulator(emulator *chip8.Chip8) []byte {
	width := 64
	height := 32
	// Create a slice to hold the RGBA values
	// Length = width * height * 4 (4 values per pixel: R, G, B, A)
	rgbaArray := make([]byte, width*height*4)
	gfx := emulator.Gfx
	// Fill the array with random RGBA values
	for i := range len(gfx) {
		if gfx[i] == 0 {
			rgbaArray[i*4] = 0
			rgbaArray[i*4+1] = 0
			rgbaArray[i*4+2] = 0
			rgbaArray[i*4+3] = 255
		} else {
			rgbaArray[i*4] = 255
			rgbaArray[i*4+1] = 255
			rgbaArray[i*4+2] = 255
			rgbaArray[i*4+3] = 255
		}

	}
	return rgbaArray

}

func (g *Game) setBeepSoundEffect() error {
	s, err := mp3.DecodeF32(bytes.NewReader(defaultBeep))
	if err != nil {
		return err
	}
	audioBytes, err := io.ReadAll(s)
	if err != nil {
		return err
	}
	g.beepSoundEffectData = audioBytes
	return nil
}

func (g *Game) playBeepSoundEffectIfNeeded() {
	if g.emulator.PlayBeepSoundEffectFlag {
		beepSePlayer := g.audioContext.NewPlayerF32FromBytes(g.beepSoundEffectData)
		beepSePlayer.Play()
		g.emulator.PlayBeepSoundEffectFlag = false
	}
}
