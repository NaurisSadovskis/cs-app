package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("  n\t 2n (base 2)\t2n (base 16)\n")
	fmt.Printf("%3d\t%12d\t0x%x\n", 5, int(math.Pow(2, 5)), int(math.Pow(2, 5)))
	fmt.Printf("%3d\t%12d\t0x%x\n", 23, int(math.Pow(2, 23)), int(math.Pow(2, 23)))
	fmt.Printf("%3d\t%12d\t0x%x\n", 32, int(math.Pow(2, 32)), int(math.Pow(2, 32)))
	fmt.Printf("%3d\t%12d\t0x%x\n", 13, int(math.Pow(2, 13)), int(math.Pow(2, 13)))
	fmt.Printf("%3d\t%12d\t0x%x\n", 12, int(math.Pow(2, 12)), int(math.Pow(2, 12)))
	fmt.Printf("%3d\t%12d\t0x%x\n", 6, int(math.Pow(2, 6)), int(math.Pow(2, 6)))
	fmt.Printf("%3d\t%12d\t0x%x\n", 12, int(math.Pow(2, 12)), int(math.Pow(2, 12)))
}

/*
Output:
  n	 2n (base 2)	2n (base 16)
  5	          32	0x20
 23	     8388608	0x800000
 32	  4294967296	0x100000000
 13	        8192	0x2000
 12	        4096	0x1000
  6	          64	0x40
 12	        4096	0x1000
*/
