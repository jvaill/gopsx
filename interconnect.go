package main

import (
    "fmt"
    "log"
)

type Interconnect struct {
    bios *BIOS
    biosRange *Range
}

func NewInterconnect(bios *BIOS) *Interconnect {
    biosRange := NewRange(0xbfc00000, 512 * 1024)
    return &Interconnect{bios: bios, biosRange: biosRange}
}

func (ic *Interconnect) read32(addr uint32) *uint32 {
    offset := ic.biosRange.contains(addr)
    if offset != nil {
        val := ic.bios.read32(*offset)
        return &val
    }

    log.Fatal(fmt.Sprintf("unknown read in interconnect at address %x", addr))
    return nil
}
