package chip8

import (
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
		chip.memory[i] = 0 // TODO: Get the chip 8 fontset

	}

	// Reset timers
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
	chip.opcode = uint16(chip.memory[chip.pc]<<8) | uint16(chip.memory[chip.pc+1])
	switch chip.opcode & 0xF000 {
	// perform opcode translation here
	}

}
