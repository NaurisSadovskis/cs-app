package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Float -> Hex conversion differs in Go, so not getting desired results
	// Switch to hex instead
	var int1 int64 = 2607352
	binaryStr := strconv.FormatInt(int1, 2)

	hex1 := "4A1F23E0"
	ui, _ := strconv.ParseUint(hex1, 16, 64)
	binaryStr2 := fmt.Sprintf("%016b", ui)

	// Add leading 0's to shorter binary
	for len(binaryStr) < len(binaryStr2) {
		binaryStr = "0" + binaryStr
	}

	fmt.Println("Hex to binary representations:")
	fmt.Println(binaryStr)
	fmt.Println(binaryStr2)

	// I know already the result
	fmt.Println("Shifting 2nd string by 2 values to the right:")
	fmt.Println(binaryStr)
	fmt.Println("00" + binaryStr2)

	fmt.Println("We can now count the matching bits above!")
}

/*
Hex to binary representations:
0000000001001111100100011111000
1001010000111110010001111100000
Shifting 2nd string by 2 values to the right:
0000000001001111100100011111000
001001010000111110010001111100000
We can now count the matching bits above!

*/
