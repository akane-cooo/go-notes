package ch01

import "fmt"

func SwitchBasic() {
	color := "yellow"
	// in Go, switch-case does not fallthrough automatically
	// so break/return is not necessary in the body of each case
	switch color {
	case "yellow":
		fmt.Println("Yellow")
	case "red":
		fmt.Println("Red")
	}
}

func double(x int) int {
	return x * 2
}

func SwitchExpression() {
	val := 10
	// evaluate expression on switch
	switch double(val) {
	case 0, 10:
		fmt.Println("almost there")
	case 20:
		fmt.Printf("find the hidden value %d\n", val)
	default:
		fmt.Println("incorrect value")
	}

	// evaluate expression in cases
	// this simulates if-else conditions
	switch {
	case double(val) > 20:
		fmt.Println("Value too large")
	case double(val) == 20:
		fmt.Println("Found correct value")
		// could have more/different evaluation here
	}
}

func Fallthrough(balance float32) {
	switch balance {
	case 100:
		fmt.Println("Good steak")
		fallthrough
	case 50:
		fmt.Println("Deserts")
		fallthrough
	case 20:
		fmt.Println("A cup of latte") // stop fallthrough from this case
	case 10:
		fmt.Println("Americano")
	default:
		fmt.Println("Insufficient funds")
	}
}

func SwitchRun() {
	Fallthrough(100)
}
