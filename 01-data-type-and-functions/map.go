package ch01

import (
	"fmt"
	"strings"
)

func MapKey() {
	// the key in the map should be hashable/comparable, cannot be
	// function, another map, slice
	type function func()

	var mapA map[int]function
	fmt.Printf("%+v\n", mapA)

	// var mapB map[function]int

	// struct can be used as the key only if all of its fileds are hashable
	type Key struct {
		ID          int
		Description string
	}
	mapB := map[Key]int{{ID: 1, Description: "placeholder test"}: 10}
	fmt.Printf("%+v\n", mapB)
}

func mapValues() {
	// map init
	var count map[string]int
	count = make(map[string]int)
	// recommend: initialize with length, avoid reallocation
	count = make(map[string]int, 5)
	count = map[string]int{"hello": 10, "nice": 2, "world": 12, "date": 1, "go": 13}

	// get value
	fmt.Printf("count of \"hello\": %d\n", count["hello"])
	if value, exists := count["typescript"]; exists {
		fmt.Printf("count of \"typescript\": %d\n", value) // default value
	}

	fmt.Printf("count of \"nice\": %d\n", count["nice"]) // default value
	delete(count, "nice")
	if value, exists := count["nice"]; exists {
		fmt.Printf("count of \"nice\": %d\n", value) // default value
	}
	count["python"] = 100

	// map has length, but not capacity
	fmt.Printf("len(count) = %d\n", len(count))

	// iterations. go creator make iteration sequence different on purpose
	for key, value := range count {
		fmt.Printf("%s: %d | ", key, value)
		value = 10
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))

	for key, value := range count {
		fmt.Printf("%s: %d | ", key, value)
		count[key] = 10
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))

	for key, value := range count {
		fmt.Printf("%s: %d | ", key, value)
		count[key] = 10
	}
	fmt.Println()
}

func MapRun() {
	// MapKey()
	mapValues()
}
