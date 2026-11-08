package ch01

import "fmt"

func GotoBasic() {
	var i int = 5
MY_LABEL: // labels, normally use use uppercase for label name
	i *= 2
	fmt.Printf("i: %d\n", i)
	if i >= 100 {
		return
	}
	goto MY_LABEL
}

func GotoIf() {
	var i int = 4
	if i%2 == 0 {
		goto L1
	} else {
		goto L2
	}
	// label is not function or condition, it will
	// keep running all of the following codes until termination condition fullfiled
L1:
	i += 3
L2:
	i += 4
	fmt.Printf("i: %d\n", i)
}

func ContinueLabel() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("Checking line %d\n", i)
	L3:
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("Checking column %d\n", j)
				if j%3 == 0 {
					// goto L1 // this will retart L1 for loop
					// break L1 // this will break the whole L1 loop
					continue L1 // this will continue L1 for loop
				} else if j%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}
	}
}

func GotoRun() {
	GotoBasic()
	GotoIf()
	ContinueLabel()
}
