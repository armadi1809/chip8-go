package chip8

import (
	"fmt"
	"log"
	"os"
)

type Chip8 struct {
	opcode      uint16
	memory      [4096]byte
	V           [16]byte
	I           uint16
	pc          uint16
	gfx         [64 * 32]byte
	delay_timer byte
	sound_timer byte
	stack       [16]uint16
	sp          uint16
	key         byte
}

var fontset []uint8 = []uint8{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

func New() *Chip8 {
	return &Chip8{
		opcode: 0,
		pc:     0x200,
		I:      0,
		sp:     0,
	}
}

func (chip *Chip8) Initialize() {
	// Clear display
	// Clear stack
	// Clear registers V0-VF
	// Clear memory

	// Load fontset
	for i := range 80 {
		chip.memory[i] = fontset[i] // TODO: Get the chip 8 fontset

	}

	// Reset timers
	chip.delay_timer = 0
	chip.sound_timer = 0
}

func (chip *Chip8) LoadProgram(path string) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("an error ocurred when reading the source code of the provided path %v", err)
	}

	for i := range len(buffer) {
		chip.memory[i+512] = buffer[i]
	}
}

func (chip *Chip8) EmulateCycle() {
	chip.opcode = uint16(chip.memory[chip.pc])<<8 | uint16(chip.memory[chip.pc+1])
	switch chip.opcode & 0xF000 {
	// perform opcode translation here
	case 0xA000:
		chip.I = chip.opcode & 0x0FFF
		chip.pc += 2
	case 0x2000:
		chip.stack[chip.sp] = chip.pc
		chip.sp++
		chip.pc = chip.opcode & 0x0FFF
	default:
		fmt.Printf("unkown opcode 0x%X\n", chip.opcode)
	}

	if chip.delay_timer > 0 {
		chip.delay_timer--
	}
	if chip.sound_timer > 0 {
		if chip.sound_timer == 1 {
			fmt.Printf("SIMULATING SOUND: BEEP\n")
		}
		chip.sound_timer--
	}

}
