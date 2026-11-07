package ch01

import "fmt"

/*
generally speaking, golang does not do "inheritance"
but we can still achieve the similar behavior with the following method
*/

type Copyright struct {
	ID,
	Country string
}

type Book struct {
	ID,
	Name string
	Copyright
}

func InheritRun() {
	copywrite := Copyright{"cd456", "us"}
	book := Book{"abc123", "Farland", copywrite}

	// we can retrieve the country value from the book instance
	fmt.Println("country", book.Country)
	// both copywrite and book have ID field, so if we want to get copywrite's id
	// we need to specify its parent struct
	fmt.Println("book id", book.ID)
	fmt.Println("copywright id", book.Copyright.ID)
}
