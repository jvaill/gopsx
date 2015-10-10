package main

type CPU struct {
    pc uint32
}

func NewCpu() *CPU {
    return &CPU{pc: 0xbfc00000}
}
