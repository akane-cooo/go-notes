package ch01

import (
	"fmt"
	"time"
)

func deferBasic() {
	fmt.Println("A")
	defer fmt.Println(1)
	fmt.Println("B")
	defer fmt.Println(2)
	fmt.Println("C")
	defer fmt.Println(3) // last registered, first executed before function complete
}

func deferWithParam() (val int) {
	val = 10

	defer func() {
		fmt.Println("first defer", val)
	}() // 1, go just hold this function

	defer func(val int) {
		fmt.Println("second defer", val)
	}(val) // 10, function hold, but param registered

	defer fmt.Println("last defer", val) // 10, direct register
	return 1
}

func deferBenchmark() {
	startT := time.Now()
	// in this method, we just evaluate the expression immediately, and print the result at the end
	defer fmt.Printf("direct defer - time spent: %.2f\n", time.Since(startT).Seconds())
	// in this method, defer registers the func without evaluating anything
	defer func() {
		fmt.Printf("func defer - time spent: %.2f\n", time.Since(startT).Seconds())
	}()
	time.Sleep(1 * time.Second)
}

func DeferRun() {
	deferBasic()
	deferWithParam()
	deferBenchmark()
}
