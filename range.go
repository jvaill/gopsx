package main

type Range struct {
    start uint32
    length uint32
}

func NewRange(start uint32, length uint32) *Range {
    return &Range{start: start, length: length}
}

func (r Range) contains(addr uint32) *uint32 {
    if (addr >= r.start && addr < r.start + r.length) {
        offset := addr - r.start
        return &offset
    }
    return nil
}
