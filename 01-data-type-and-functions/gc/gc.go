package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"runtime/trace"
	"sync"
	"time"
)

// var (
// 	arr [1000]int                     // value type (including array) is assigned on the global data segment
// 	slc []int     = make([]int, 1000) // reference-type global variable will be assigned on the heap
// )

// go run -gcflags=-m ./01-data-type-and-functions/gc

func StackHeap() {
	var brr [1000]byte        // function's in/out/variables will be on stack when it is small
	var crr [128 << 10]byte   // 128KB - max amount on stack for primitive (array) type
	var drr [128<<10 + 1]byte // over 128KB, will escape to the heap

	err := make([]byte, 64<<10)   // reference-type, max 64KB on stack
	frr := make([]byte, 64<<10+1) // escape to heap

	_ = brr
	_ = crr
	_ = drr
	_ = err
	_ = frr
}

// =========== Followings are for the GC simulation ==========

func SimilateGC() {
	f, err := os.Create("./data/trace.out")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer trace.Stop()

	debug.SetGCPercent(100)

	numTasks := 1000
	inputChan := make(chan int, numTasks)
	outputChan := make(chan int, numTasks)

	var wg sync.WaitGroup

	// create workgroup and start pushing tasks into input channel
	numWorkers := 5
	wg.Add(numWorkers)
	for range numWorkers {
		go worker(inputChan, outputChan, &wg)
	}
	for i := range numTasks {
		inputChan <- i
	}
	close(inputChan)

	go func() {
		wg.Wait()
		close(outputChan)
	}()
	for val := range outputChan {
		_ = val
	}
	fmt.Println("Done")
}

func worker(inputChan <-chan int, outputChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for input := range inputChan {
		result := memIntensiveTask(input)
		outputChan <- result
	}
}

func memIntensiveTask(input int) int {
	largeMemSize := 1<<20 + 1
	localSlc := make([]int, largeMemSize)
	for i := range largeMemSize {
		localSlc[i] = input + i
	}
	// simulate latency
	time.Sleep(3 * time.Millisecond)
	result := 0
	for _, val := range localSlc {
		result += val
	}
	return result
}

func main() {
	// StackHeap()
	// go tool trace data/trace.out
	SimilateGC()
}
