package main

const biosRomPath = "./SCPH1001.BIN"

func main() {
    cpu := NewCpu()
    bios := NewBios(biosRomPath)

    _, _ = cpu, bios
}
