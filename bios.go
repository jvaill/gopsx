package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "crypto/md5"
)

// BIOS ROMs are always 512kb in length.
const biosRomLength int = 1024 * 512

type BIOS struct {
    data []byte
}

func NewBios(biosRomPath string) *BIOS {
    b, err := ioutil.ReadFile(biosRomPath)
    if err != nil {
        log.Fatal("could not read bios: ", err)
    }

    if len(b) != biosRomLength {
        log.Fatal(fmt.Sprintf("expected bios to be %d bytes, got %d bytes", biosRomLength, len(b)))
    }

    log.Print(fmt.Sprintf("bios md5 checksum is 0x%x", md5.Sum(b)))
    return &BIOS{data: b}
}
