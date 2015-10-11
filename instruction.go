package main

import (
    "fmt"
)

// TODO: What are the performance implications of creating a struct for each
//       new instruction?

type Instruction struct {
    opcode uint32
}

func NewInstruction(opcode uint32) *Instruction {
    return &Instruction{opcode: opcode}
}

func (instruction *Instruction) String() string {
    return fmt.Sprintf("0x%x", instruction.opcode)
}

func (instruction *Instruction) primaryOpcode() uint8 {
    b := uint8(instruction.opcode >> 26)
    return b
}

func (instruction *Instruction) sourceRegister() uint8 {
    b := uint8((instruction.opcode >> 21) & 0x1F)
    return b
}

func (instruction *Instruction) targetRegister() uint8 {
    b := uint8((instruction.opcode >> 16) & 0x1F)
    return b
}

func (instruction *Instruction) imm16() uint16 {
    val := uint16(instruction.opcode & 0xFFFF)
    return val
}
