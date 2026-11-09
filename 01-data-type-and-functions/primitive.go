package ch01

import (
	"fmt"
	"math"
)

func PrimitiveConversions() {
	var val int16 = 128 // 1000 0000
	val8 := int8(val)   // 1000 0000
	fmt.Printf("int16: %d; int8: %d\n", val, val8)

	var uval64 uint64 = math.MaxUint64
	val64 := int64(uval64)
	fmt.Printf("uint64: %d; int64: %d\n", uval64, val64)

	uval32 := uint32(uval64)
	fmt.Printf("uint64: %d; uint32: %d\n", uval64, uval32)

	// char
	cVal := int('k')
	fmt.Printf("%d\n", cVal)

	// bool and int cannot be converted

	// byte - int
	by := byte(cVal)
	fmt.Printf("byte: %b\n", by)
	i := int(by)
	fmt.Printf("int: %d\n", i)

	// float int
	ft := float32(13.567)
	fmt.Printf("float: %f\n", ft)
	i = int(ft)
	fmt.Printf("int: %d\n", i)
}

func PrimitiveRun() {
	PrimitiveConversions()
}
