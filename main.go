package main

const biosRomPath = "./SCPH1001.BIN"

func main() {
    bios := NewBios(biosRomPath)
    interconnect := NewInterconnect(bios)
    cpu := NewCpu(interconnect)

    cpu.run()
}
