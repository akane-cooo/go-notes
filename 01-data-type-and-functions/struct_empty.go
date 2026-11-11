package ch01

import (
	"fmt"
	"reflect"
	"time"
)

type Empty struct{}

func AddressAndMemory() {
	var a Empty
	var b Empty
	c := Empty{}
	fmt.Printf("addresses a: %p, b: %p, c:%p\n", &a, &b, &c)
	fmt.Printf("vals a: %v, b: %v, c:%v\n", a, b, c)
	fmt.Printf("sizes a: %d, b: %d, c:%d\n",
		reflect.TypeOf(a).Size(),
		reflect.TypeOf(b).Size(),
		reflect.TypeOf(c).Size())
}

func SetImpl() {
	set := map[string]struct{}{}
	set["a"] = struct{}{}
	set["b"] = struct{}{}
	if _, exists := set["c"]; !exists {
		fmt.Println("c not in the set")
	}
}

func MockBlockingProcess() {
	timeFormat := "15:04:05"
	blocker := make(chan struct{})
	go func() {
		fmt.Printf("curr time: %s\n", time.Now().Format(timeFormat))
		time.Sleep(2 * time.Second)
		blocker <- struct{}{}
	}()
	<-blocker
	fmt.Printf("curr time: %s\n", time.Now().Format(timeFormat))
}

func EmptyStructRun() {
	AddressAndMemory()
	SetImpl()
	MockBlockingProcess()
}
