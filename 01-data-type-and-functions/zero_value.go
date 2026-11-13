package ch01

import "fmt"

func ZeroValueRun() {
	// primitives
	fmt.Println("default value for primitives")
	var i int
	fmt.Printf("int: %v\n", i)
	var f float32
	fmt.Printf("float32: %v\n", f)
	var b byte
	fmt.Printf("byte: %v\n", b)
	var bl bool
	fmt.Printf("bool: %v\n", bl)
	var s string
	fmt.Printf("string: %q\n", s)
	var p *int
	fmt.Printf("address: %v\n", p)

	fmt.Println("")
	fmt.Println("default value for struct and reference types")
	type ZeroStruct struct {
		ID   int
		Name string
	}
	var u ZeroStruct
	fmt.Printf("struct ZeroStruct: %+v\n", u)
	var err error
	fmt.Printf("error: %v\n", err)
	var arr [3]int
	fmt.Printf("[3]int: %v\n", arr)
	var slc []string
	fmt.Printf("[]string: %v\n", slc)
	var mp map[int]byte
	fmt.Printf("map[int]byte: %v\n", mp)
	var ch chan float32
	fmt.Printf("chan float32: %v\n", ch)

	// channels
	chs := make(chan int, 10)
	close(chs)
	val, ok := <-chs
	fmt.Printf("closed channel value: %v, ok: %t\n", val, ok)
}
