package main

import "fmt"

// Added additional int and string print
func main() {
	fmt.Printf("hex\tint\tstring\n")
	fmt.Printf("%x\t%d\t%v\n", 0x605c+0x5, int(0x605c+0x5), string(0x605c+0x5))
	fmt.Printf("%x\t%d\t%v\n", 0x605c-0x20, int(0x605c-0x20), string(0x605c-0x20))
	fmt.Printf("%x\t%d\t%v\n", 0x605c+32, int(0x605c+32), string(0x605c+32))
	fmt.Printf("%x\t%d\t%v\n", 0x60fa-0x605c, int(0x60fa-0x605c), string(0x60fa-0x605c))

}

/*
Output:
hex	int	string
6061	24673	恡
603c	24636	怼
607c	24700	恼
9e		158	
*/
