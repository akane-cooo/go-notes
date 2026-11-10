package ch01

import "fmt"

/*
* panic may happen when system initialization failed
* or index out of range, or zero division
*
* when panic happend:
* 1. execute registered defer
* 2. trace and print errors
* 3. end go. (by calling os.Exit(2))
* */

func panicBasic() {
	data := []int{1, 2, 3}
	defer func() {
		fmt.Println("data created")
	}()

	defer func() {
		_ = data[4]
		fmt.Println("data retrieved")
	}()

	defer func() {
		fmt.Println("last registered defer")
	}()
}

func recoverBasic() {
	data := []int{1, 2, 3}
	defer func() {
		// recover has to be in the defer func()
		if panicInfo := recover(); panicInfo != nil {
			fmt.Printf("panic in function recoverBasic: %v\n", panicInfo)
		}
	}()
	_ = data[3]
	// the recover wont work here, because the defer has not been registered
	// defer func() {
	// 	// recover has to be in the defer func()
	// 	if panicInfo := recover(); panicInfo != nil {
	// 		fmt.Printf("panic in function recoverBasic: %v\n", panicInfo)
	// 	}
	// }()
}

func PanicRun() {
	// panicBasic()
	recoverBasic()
}
