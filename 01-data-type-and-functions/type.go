package ch01

import "fmt"

type (
	Integer int   // int and Integer can covert to eachother, because the data types are the same
	Int     = int // Int and int are equivalent
)

func TypeRun() {
	var a Integer
	var b Int
	var original int

	a = 10
	b = 10

	original = int(a)
	original = b
	fmt.Println(original)
}
