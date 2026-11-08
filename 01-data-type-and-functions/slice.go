package ch01

import (
	"fmt"
	"strings"
)

func SliceInit() {
	// 1D slice
	var s []int // declaration, len=0 cap=0
	fmt.Printf("len %d, cap %d\n", len(s), cap(s))
	s = []int{} // initialization, len=0 cap=0
	fmt.Printf("len %d, cap %d\n", len(s), cap(s))
	s = make([]int, 3) // make, len=0 cap=3
	fmt.Printf("len %d, cap %d\n", len(s), cap(s))
	s = make([]int, 3, 5) // make, len=0 cap=3
	fmt.Printf("len %d, cap %d\n", len(s), cap(s))
	s = []int{1, 2, 3, 4, 5} // initialize with values, len=cap=5
	fmt.Printf("len %d, cap %d\n", len(s), cap(s))

	// 2D slice
	// child slice in 2d slice can have different length and capacity
	s2d := [][]int{
		make([]int, 2, 10),
		make([]int, 3, 4),
	}
	fmt.Printf("2D slice, len %d, cap %d\n", len(s2d), cap(s2d))
	fmt.Printf("s2d[0], len %d, cap %d\n", len(s2d[0]), cap(s2d[0]))
	fmt.Printf("s2d[1], len %d, cap %d\n", len(s2d[1]), cap(s2d[1]))
	fmt.Println(strings.Repeat("=", 20))
}

func SliceExpansion() {
	prevCap := 6
	s := make([]int, 3, prevCap)
	for i := range 100 {
		s = append(s, i)
		currCap := cap(s)
		if currCap != prevCap {
			fmt.Printf(
				"Slice expanded, prev %d, curr %d, addr %p\n",
				prevCap,
				currCap,
				s,
			)
			prevCap = currCap
		}
	}
}

func SliceRun() {
	SliceInit()
	SliceExpansion()
}
