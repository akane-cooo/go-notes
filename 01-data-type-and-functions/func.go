package ch01

import "fmt"

func Return1(a, b int) int {
	c := a + b
	return c
}

// Return2 the return variable name has been initialized
func Return2(a, b int) (c int) {
	c = a + b
	return // the return statement is still required
}

// Return3 in go, the function will make a copy of input params
// therefore modifying values inside the function won't
// affect the param's original value
// if the input param is larger, copying params will inevitably increase the
// program's memory usage
func Return3(a, b int) {
	a = a + b
}

// Return4 if we pass in the address of the param
// we can modify the original value
func Return4(a, b *int) {
	*a = *a + *b
	*b = 276
}

// variadic params

func sum(start int, arr ...int) (ret int) {
	// arr here is actually a slice
	ret = start
	fmt.Printf("arr len=%d, cap=%d\n", len(arr), cap(arr))
	for _, val := range arr {
		ret += val
	}
	return
}

func variadicParams() {
	vals := []int{10, 12, 6, 9}
	result := sum(10, vals...)
	fmt.Println("result: ", result)
}

func FuncRun() {
	a, b := 16, 27
	fmt.Printf("Original: a=%d, b=%d\n", a, b)
	Return3(a, b)
	fmt.Printf("After Return3: a=%d, b=%d\n", a, b)
	Return4(&a, &b)
	fmt.Printf("After Return3: a=%d, b=%d\n", a, b)

	variadicParams()
}
