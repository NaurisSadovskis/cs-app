package main

import (
	"encoding/binary"
	"fmt"
)

func showBytes(start uint32, lenght int) {
	bs := make([]byte, 4)
	ls := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, start)
	binary.LittleEndian.PutUint32(ls, start)
	fmt.Printf("Big Endian: %d\tLittle Endian: %d\n", bs[0:lenght], ls[0:lenght])
}

func main() {
	var a uint32 = 0x12345678
	showBytes(a, 1)
	showBytes(a, 2)
	showBytes(a, 3)
}

/*
Output:
Big Endian: [18]	Little Endian: [120]
Big Endian: [18 52]	Little Endian: [120 86]
Big Endian: [18 52 86]	Little Endian: [120 86 52]
*/
