package ch01

import "fmt"

func accumulate(vals ...interface{}) (ret int) {
	for _, val := range vals {
		// check the data type of the input val
		switch v := val.(type) { // .type can only go after switch
		case int:
			ret += v
		case float32:
			ret += int(v)
		case float64:
			ret += int(v)
		default:
			fmt.Printf("Unaccumulatable value %[1]T, %[1]v\n", val) // specify the which value to be formatted
		}
	}
	return
}

func anyMap() {
	records := make(map[any]any)
	records[1] = "1"
	records["hello"] = 124

	// type assertion
	record := records["hello"]
	if val, ok := record.(string); ok {
		fmt.Printf("string %s\n", val)
	} else if val, ok := record.(int); ok {
		fmt.Printf("int %d\n", val)
	}
}

func accumulateValues() {
	testSlice := make([]any, 0, 5)
	testSlice = append(testSlice, 10)
	testSlice = append(testSlice, float32(2.64))
	testSlice = append(testSlice, float64(1.43))
	testSlice = append(testSlice, "asb")
	testSlice = append(testSlice, struct {
		Name   string
		Weight int
	}{"milk", 20})
	result := accumulate(testSlice...)
	fmt.Printf("accumulate retulr: %d\n", result)
}

func AnyRun() {
	accumulateValues()
	anyMap()
}
