package ch01

import "fmt"

const PI = 1

func ScopeRun() {
	// "{}" defines the scope of variables
	{
		const PI = 3.14 // float
		const ScopedE = 20
		fmt.Printf("PI: %f, ScopedE: %d\n", PI, ScopedE) // this should be 3.14, closest first
	}
	{
		// fmt.Println(ScopedE) this doesn't work because ScopedE is not in the same scope
		fmt.Printf("PI: %d\n", PI) // this should be 3.14, closest first
	}
}
