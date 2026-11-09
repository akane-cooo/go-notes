// Package ch01 data type and functions
package ch01

import (
	"fmt"
	"strings"
)

func Array1D() {
	arr1 := [5]int{} // 5 int values, default to 0
	arr2 := [5]int{12, 32, 65, 12, 90}
	arr3 := [5]int{0: 10, 4: 20}
	arr4 := [...]int{45, 12}
	arr5 := [...]struct {
		Name string
		Age  int
	}{
		{"Lisa", 20}, {"Nick", 12},
	}

	fmt.Printf("arr1=%+#v\n", arr1)
	fmt.Printf("arr2=%+#v\n", arr2)
	fmt.Printf("arr3=%+#v\n", arr3)
	fmt.Printf("arr4=%+#v\n", arr4)
	fmt.Printf("arr5=%+#v\n", arr5)
	fmt.Println(strings.Repeat("=", 20))
}

func Array2D() {
	arr1 := [5][3]int{{2: 10}, {10, 20, 34}, 3: {13, 54, 9}}
	arr2 := [...][2]int{{1, 2}, {3, 5}}

	fmt.Printf("arr1=%+v\n", arr1)
	// for array, capacity and length are the same
	fmt.Printf("arr2 length: %d\n", len(arr2))
	fmt.Printf("arr2 cap: %d\n", cap(arr2)) // capacity
}

func ArrayInFunction() {
	arr := [4]int{74, 10, 68, 33}
	fmt.Printf("val of arr[2]: %d\n", arr[2])
	updateArr1(arr)
	fmt.Printf("val of arr[2]: %d\n", arr[2])
	updateArr2(&arr)
	fmt.Printf("val of arr[2]: %d\n", arr[2])
}

func updateArr1(arr [4]int) {
	arr[2] = 1234
}

func updateArr2(arr *[4]int) {
	arr[2] = 1234
}

func ArrayRun() {
	Array1D()
	Array2D()
	ArrayInFunction()
}
