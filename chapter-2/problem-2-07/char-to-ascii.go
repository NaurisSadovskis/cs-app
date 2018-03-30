package main

import "fmt"

func main() {
	for _, letter := range "mnopqr" {
		fmt.Printf("%x ", letter)
	}
	fmt.Println()
}

/*
6d 6e 6f 70 71 72
*/
