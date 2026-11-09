package ch01

import (
	"fmt"
	"slices"
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

func SliceSharedMemory() {
	arr := [4]int{2, 3, 58, 2}
	brr := arr[2:3]
	// the following two addresses should be the same
	fmt.Printf("addr of arr[2]: %p\n", &arr[2])
	fmt.Printf("addr of brr: %p\n", brr)
	// this append modifies the value right after arr[2]
	brr = append(brr, 65)
	// the this append exceeds [4]int's capacity
	// from this point, the brr has been allocated to a new address
	brr = append(brr, 75)
	brr[1] = 165
	fmt.Printf("%+v", arr)
	fmt.Printf("%+v", brr)
}

func SliceSort() {
	// regular sort
	arr := []int{498, 24, 45, 121, 0, 12297}
	slices.Sort(arr)
	fmt.Printf("%v\n", arr)
	slices.SortFunc(arr, func(a, b int) int {
		return b - a
	})
	fmt.Printf("%v\n", arr)

	// struct sort by field
	type Plant struct {
		Name   string
		Height float32
	}
	plants := []Plant{
		{Name: "treeA", Height: 12.5},
		{Name: "treeB", Height: 16.5},
		{Name: "flower", Height: 4.0},
	}
	slices.SortFunc(plants, func(a, b Plant) int {
		if b.Height > a.Height {
			return 1
		} else if b.Height < a.Height {
			return -1
		} else {
			return 0
		}
	})
	fmt.Printf("%v\n", plants)
}

func SliceCopyAndEqual() {
	arr := []int{2, 4, 7, 10, 22, 0, 1}
	brr := make([]int, len(arr))
	copy(brr, arr)
	fmt.Printf("arr: %v; brr: %v\n", arr, brr)
	fmt.Printf("arr == brr: %t\n", slices.Equal(arr, brr))
	crr := make([]int, len(arr)-2)
	copy(crr, arr)
	fmt.Printf("arr: %v; crr: %v\n", arr, crr)
	fmt.Printf("arr == crr: %t\n", slices.Equal(arr, crr))
}

func SliceRun() {
	// SliceInit()
	// SliceExpansion()
	// SliceSharedMemory()
	// SliceSort()
	SliceCopyAndEqual()
}
