package ch01

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("data not found")
	ErrTooLarge = errors.New("input too large")
	ErrNegative = errors.New("negative value")
)

func funcA(a int) (int, error) {
	if a < 0 {
		return 0, MyError{Name: "negative value", Code: 0, Desc: ErrNegative.Error()}
	}
	result, err := funcB(a)
	if err != nil {
		var myErr MyError
		switch {
		case errors.Is(err, ErrNotFound): // check is the error a wrapper of the other(base one)
			myErr = MyError{Name: "negative value", Code: 1}
		case errors.Is(err, ErrTooLarge):
			myErr = MyError{Name: "large value", Code: 2}
		}
		myErr.Desc = fmt.Errorf("funcA -> %w", err).Error()
		return 0, myErr
	}
	return result, nil
}

func funcB(b int) (int, error) {
	if b > 10 {
		return 0, ErrTooLarge
	}
	result, err := funcC(b)
	if err != nil {
		return 0, fmt.Errorf("funcB -> %w", err)
	}
	return result, nil
}

func funcC(c int) (int, error) {
	if c == 0 {
		return 0, ErrNotFound
	}
	return c * 2, nil
}

func ErrTraceRun() {
	result, err := funcA(20)
	var myErr MyError
	if err != nil {
		ok := errors.As(err, &myErr) // type assertion
		if ok {
			fmt.Println(myErr.Desc)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(result)
	}
}
