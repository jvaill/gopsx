package main

import (
    "fmt"
    "log"
)

type CPU struct {
    interconnect *Interconnect
    pc uint32
}

func NewCpu(interconnect *Interconnect) *CPU {
    return &CPU{interconnect: interconnect, pc: 0xbfc00000}
}

func (cpu CPU) read32(addr uint32) uint32 {
    return *cpu.interconnect.read32(addr)
}

func (cpu CPU) decodeAndExecute(instruction uint32) {
    log.Fatal(fmt.Sprintf("unhandled instruction 0x%x", instruction))
}

func (cpu CPU) run() {
    for {
        instruction := cpu.read32(cpu.pc)
        cpu.decodeAndExecute(instruction)
        cpu.pc += 4
    }
}
