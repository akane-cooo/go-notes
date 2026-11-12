package ch01

import (
	"errors"
	"fmt"
)

var ErrZeroDivision = errors.New("zero division") // start with Err is an error naming customs

type MyError struct {
	Name string
	Code int
	Desc string
}

func (e MyError) String() string {
	return fmt.Sprintf("%s - %s", e.Name, e.Desc)
}

func (e MyError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Desc)
}

// by customs, we put error as the last return value
func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, ErrZeroDivision
	} else {
		return a / b, nil
	}
}

func mockResponse() MyError {
	return MyError{Name: "Response Error", Code: 404, Desc: "Resource not found"}
}

func ErrorRun() {
	result, err := divide(10, 22)
	if err != nil {
		fmt.Printf("division failed: %s\n", err.Error())
	} else {
		fmt.Printf("result: %.4f\n", result)
	}

	myErr := mockResponse()
	fmt.Printf("division failed: %s\n", myErr) // if the customized has Error and String method, system will use prioritize Error
}
