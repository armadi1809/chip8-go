package widgets

import (
	"bytes"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/text/language"
)

type ComboBox struct {
	X, Y             float64   // Position of the combo box
	Width, Height    float64   // Dimensions of the combo box
	Options          []string  // List of options
	SelectedIndex    int       // Index of the currently selected option
	IsOpen           bool      // Whether the dropdown is open
	Font             text.Face // Font for rendering text
	lastClickTime    time.Time
	debounceDuration time.Duration
}

func NewComboBox(x, y, width, height float64, options []string) *ComboBox {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	f := &text.GoTextFace{
		Source:   s,
		Size:     16,
		Language: language.English,
	}
	return &ComboBox{
		X:                x,
		Y:                y,
		Width:            width,
		Height:           height,
		Options:          options,
		SelectedIndex:    0,
		IsOpen:           false,
		Font:             f, // Use a basic font
		lastClickTime:    time.Time{},
		debounceDuration: 250 * time.Millisecond,
	}
}

func (cb *ComboBox) Draw(screen *ebiten.Image) {
	// Draw the main box
	boxColor := color.RGBA{200, 200, 200, 255} // Light gray
	if cb.IsOpen {
		boxColor = color.RGBA{180, 180, 180, 255} // Slightly darker when open
	}
	vector.DrawFilledRect(screen, float32(cb.X), float32(cb.Y), float32(cb.Width), float32(cb.Height), boxColor, true)

	// Draw the selected option
	selectedOption := cb.Options[cb.SelectedIndex]
	op := &text.DrawOptions{}

	op.ColorScale.SetA(255)
	op.ColorScale.SetR(255)
	op.ColorScale.SetG(255)
	op.ColorScale.SetB(255)

	op.GeoM.Translate(cb.X+5, cb.Y)

	text.Draw(screen, selectedOption, cb.Font, op)

	// Draw the dropdown menu if open
	if cb.IsOpen {
		for i, option := range cb.Options {
			optionY := cb.Y + float64(i+1)*cb.Height
			vector.DrawFilledRect(screen, float32(cb.X), float32(optionY), float32(cb.Width), float32(cb.Height), boxColor, true)
			op = &text.DrawOptions{}
			op.ColorScale.SetA(255)
			op.ColorScale.SetR(255)
			op.ColorScale.SetG(255)
			op.ColorScale.SetB(255)
			op.GeoM.Translate(cb.X+5, optionY)
			text.Draw(screen, option, cb.Font, op)
		}
	}
}

func (cb *ComboBox) Update() {

	currentTime := time.Now()

	// Check if enough time has passed since the last click
	if currentTime.Sub(cb.lastClickTime) < cb.debounceDuration {
		return
	}

	mouseX, mouseY := ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if float64(mouseX) >= cb.X && float64(mouseX) <= cb.X+cb.Width &&
			float64(mouseY) >= cb.Y && float64(mouseY) <= cb.Y+cb.Height {
			cb.IsOpen = !cb.IsOpen // Toggle dropdown
			cb.lastClickTime = time.Now()
		} else if cb.IsOpen { // Check if an option is clicked
			for i := range cb.Options {
				optionY := cb.Y + float64(i+1)*cb.Height
				if float64(mouseX) >= cb.X && float64(mouseX) <= cb.X+cb.Width &&
					float64(mouseY) >= optionY && float64(mouseY) <= optionY+cb.Height {
					cb.SelectedIndex = i
					cb.IsOpen = false // Close dropdown after selection
					cb.lastClickTime = time.Now()
					break
				}
			}
		}
	}
}
