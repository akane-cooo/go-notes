package ch01

import (
	"fmt"
	"strconv"
	"strings"
)

func StringRune() {
	s1 := "Hello my name is あかね"
	byteArr := []byte(s1)
	runeArr := []rune(s1)
	fmt.Printf("last byte %d\n", byteArr[len(byteArr)-1])
	fmt.Printf("last byte %q\n", byteArr[len(byteArr)-1])
	fmt.Printf("last rune %d\n", runeArr[len(runeArr)-1])
	fmt.Printf("last rune %q\n", runeArr[len(runeArr)-1])

	L := len(s1)
	// rune iteration
	for _, ele := range s1 {
		fmt.Printf("%c ", ele)
	}
	fmt.Println()
	// byte iteration
	for i := range L {
		fmt.Printf("%c ", s1[i])
	}
	fmt.Println()

	// rune split
	arr := strings.Split(s1, " ")
	fmt.Printf("%v\n", arr)
}

func StringJoin() {
	s1 := "Hello"
	s2 := "my"
	s3 := "name"
	s4 := "is"
	s5 := "あかね"

	s := s1 + " " + s2 + " " + s3 + " " + s4 + " " + s5
	fmt.Println("concactenation: ", s)
	s = fmt.Sprintf("%s %s %s %s %s", s1, s2, s3, s4, s5)
	fmt.Println("fmt.Sprintf: ", s)
	s = strings.Join([]string{s1, s2, s3, s4, s5}, " ")
	fmt.Println("strings.Join: ", s)

	// recommended for large text building, better performance
	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(" ")
	sb.WriteString(s2)
	sb.WriteString(" ")
	sb.WriteString(s3)
	sb.WriteString(" ")
	sb.WriteString(s4)
	sb.WriteString(" ")
	sb.WriteString(s5)
	merged := sb.String()
	fmt.Println("strings.Builder: ", merged)
}

func CommonOperations() {
	fmt.Printf("Length of \"s\": %d\n", len("s"))
	fmt.Printf("Length of \"あ\": %d\n", len("あ"))

	s := "hello this time i am 飛鳥 hello"
	// contains
	fmt.Printf("contains time: %t\n", strings.Contains(s, "time"))
	fmt.Printf("contains 飛 %t\n", strings.Contains(s, "飛"))
	// index
	fmt.Printf("Index of time: %d\n", strings.Index(s, "time"))
	fmt.Printf("Index of 飛 %d\n", strings.Index(s, "飛"))
	fmt.Printf("Last index of hello %d\n", strings.LastIndex(s, "hello"))
	// Prefix and suffix
	fmt.Printf("Has prefix hello %t\n", strings.HasPrefix(s, "hello"))
	fmt.Printf("Has suffix hello %t\n", strings.HasSuffix(s, "hello"))
}

func Conversions() {
	i := 8
	i64 := int64(i)

	// int to string
	s := strconv.Itoa(i)
	fmt.Printf("%s\n", s)
	s = strconv.FormatInt(i64, 10)
	fmt.Printf("%s\n", s)

	// string to int
	i, _ = strconv.Atoi(s)
	fmt.Printf("%d\n", i)
	i, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Printf("Failed to conver string to int: %v\n", err)
	} else {
		fmt.Printf("%d\n", i)
	}

	// float to string
	val := 43.12591
	s = strconv.FormatFloat(val, 'f', 2, 32) // precison, 4 bytes
	fmt.Println(s)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Conversion failed: %s\n", err)
	}
	fmt.Println(f)
}

func StringRun() {
	// StringRune()
	// StringJoin()
	// CommonOperations()
	Conversions()
}
