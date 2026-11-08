package ch01

import "fmt"

func Return1(a, b int) int {
	c := a + b
	return c
}

// the return variable name has been initialized
func Return2(a, b int) (c int) {
	c = a + b
	return // the return statement is still required
}

// in go, the function will make a copy of input params
// therefore modifying values inside the function won't
// affect the param's original value
// if the input param is larger, copying params will inevitably increase the
// program's memory usage
func Return3(a, b int) {
	a = a + b
}

// if we pass in the address of the param
// we can modify the original value
func Return4(a, b *int) {
	*a = *a + *b
	*b = 276
}

func FuncRun() {
	a, b := 16, 27
	fmt.Printf("Original: a=%d, b=%d\n", a, b)
	Return3(a, b)
	fmt.Printf("After Return3: a=%d, b=%d\n", a, b)
	Return4(&a, &b)
	fmt.Printf("After Return3: a=%d, b=%d\n", a, b)
}
