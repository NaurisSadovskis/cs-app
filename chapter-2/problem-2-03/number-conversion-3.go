package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Printf("base10\t\tbase2\tbase 16\n")
	fmt.Printf("%6d\t%13b\t0x%X\n", 0, 0, 0)
	fmt.Printf("%6d\t%13b\t0x%X\n", 158, 158, 158)
	fmt.Printf("%6d\t%13b\t0x%X\n", 78, 78, 78)
	bin1, _ := strconv.ParseInt("10101110", 2, 64)
	fmt.Printf("%6d\t%13b\t0x%X\n", bin1, bin1, bin1)
	bin2, _ := strconv.ParseInt("00111100", 2, 64)
	fmt.Printf("%6d\t%13b\t0x%X\n", bin2, bin2, bin2)
	bin3, _ := strconv.ParseInt("11110001", 2, 64)
	fmt.Printf("%6d\t%13b\t0x%X\n", bin3, bin3, bin3)
}

/*
Output:
base10		base2	base 16
     0	            0	0x0
   158	     10011110	0x9E
    78	      1001110	0x4E
   174	     10101110	0xAE
    60	       111100	0x3C
   241	     11110001	0xF1
*/
