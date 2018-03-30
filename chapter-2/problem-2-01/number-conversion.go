package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Hexadecimal to binary
	rawHex1 := "25B9D2"
	rawHex2 := "A8B3D"
	int1, _ := strconv.ParseInt(rawHex1, 16, 64)
	int2, _ := strconv.ParseInt(rawHex2, 16, 64)
	bin1 := strconv.FormatInt(int1, 2)
	bin2 := strconv.FormatInt(int2, 2)

	// Binary to hexadecimal
	rawBinary1 := "1010111001001001"
	rawBinary2 := "1100100010110110010110"
	int3, _ := strconv.ParseInt(rawBinary1, 2, 64)
	int4, _ := strconv.ParseInt(rawBinary2, 2, 64)

	fmt.Printf("Answer A:\nHexadecimal: 0x%s\nInteger: %d\nBinary: %s\n\n", rawHex1, int1, bin1)
	fmt.Printf("Answer B:\nHexadecimal: 0x%x\nInteger: %d\nBinary: %s\n\n", int3, int3, rawBinary1)
	fmt.Printf("Answer C:\nHexadecimal: 0x%s\nInteger: %d\nBinary: %s\n\n", rawHex2, int2, bin2)
	fmt.Printf("Answer D:\nHexadecimal: 0x%x\nInteger: %d\nBinary: %s\n\n", int4, int4, rawBinary2)
}
