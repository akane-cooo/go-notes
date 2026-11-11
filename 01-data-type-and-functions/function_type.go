package ch01

import "fmt"

type Compute func(a, b, c int, d bool) float32

func (Compute) Clear() {
	fmt.Println("clear")
}

type Processor struct {
	ID          string
	ComputeLoad Compute
}

func funcType() {
	f1 := func(a, b, c int, d bool) float32 {
		if d {
			return float32(a + b + c)
		}
		return float32(0.1)
	}
	f1(1, 2, 3, true)

	// the f1 doesn't have Clear method, but Compute has
	// therefore we can still use the Clear method in this case
	processor := Processor{ID: "123", ComputeLoad: f1}
	processor.ComputeLoad(1, 2, 3, true)
	processor.ComputeLoad.Clear()
}

func FuncTypeRun() {
	funcType()
}
