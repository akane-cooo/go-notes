package ch01_test

import (
	"fmt"
	"testing"

	ch01 "github.com/akane-cooo/go-notes/01-data-type-and-functions"
)

func TestSet(t *testing.T) {
	intSet := ch01.NewSet[int](3)
	intSet.Add(10)
	intSet.Add(20)
	intSet.Add(30)
	if !intSet.Has(10) || !intSet.Has(20) || !intSet.Has(30) {
		t.Fail()
	}
	if intSet.Len() != 3 {
		t.Fail()
	}
	intSet.Remove(10)
	if intSet.Has(10) {
		t.Fail()
	}

	intSet.Range(func(ele int) {
		fmt.Printf("%d ", ele)
	})
	fmt.Printf("\n")
}
