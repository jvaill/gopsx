package main

import (
    "fmt"
    "log"
)

type CPU struct {
    interconnect *Interconnect
    pc uint32
    regs [32]uint32
}

func NewCpu(interconnect *Interconnect) *CPU {
    return &CPU{interconnect: interconnect, pc: 0xbfc00000}
}

func (cpu *CPU) reg(index uint8) uint32 {
    return cpu.regs[index]
}

func (cpu *CPU) setReg(index uint8, val uint32) {
    cpu.regs[index] = val

    // Ensure that the $zero register is always 0.
    cpu.regs[0] = 0
}

func (cpu *CPU) read32(addr uint32) uint32 {
    return *cpu.interconnect.read32(addr)
}

func (cpu *CPU) ori(instruction *Instruction) {
    rs := instruction.sourceRegister()
    rt := instruction.targetRegister()
    imm := instruction.imm16()

    val := cpu.reg(rs) | uint32(imm)
    cpu.setReg(rt, val)
}

func (cpu *CPU) lui(instruction *Instruction) {
    rt := instruction.targetRegister()
    imm := instruction.imm16()

    // Low 16 bits are set to 0.
    val := uint32(imm) << 16
    cpu.setReg(rt, val)
}

func (cpu *CPU) decodeAndExecute(instruction *Instruction) {
    switch instruction.primaryOpcode() {
        default:
            log.Fatal(fmt.Sprintf(
                "unhandled instruction %s with primary opcode 0x%x",
                instruction, instruction.primaryOpcode()))
        case 0xD:
            cpu.ori(instruction)
        case 0xF:
            cpu.lui(instruction)
    }
}

func (cpu *CPU) run() {
    for {
        val := cpu.read32(cpu.pc)
        instruction := NewInstruction(val)
        cpu.decodeAndExecute(instruction)
        cpu.pc += 4
    }
}
