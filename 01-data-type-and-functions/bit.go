package ch01

import "fmt"

func BinaryFormat(val uint16) {
	var c uint16 = 1 << 15
	for range 16 {
		switch (val & c) == c {
		case true:
			fmt.Print(1)
		default:
			fmt.Print(0)
		}
		c >>= 1
	}
	fmt.Printf("\n")
	fmt.Printf("binary print: %b\n", val)
}

func BinaryRun() {
	BinaryFormat(165)
}
