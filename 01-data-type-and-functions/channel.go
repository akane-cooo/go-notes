package ch01

import "fmt"

// channel is basically a FIFO queue
func rangeChannel() {
	ch := make(chan int, 8)
	ch <- 1 // start sending data to the channel
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Printf("chan len %d, cap %d\n", len(ch), cap(ch))
	v := <-ch
	fmt.Printf("v %d\n", v)
	v = <-ch
	fmt.Printf("v %d\n", v)

	// read/write only channel
	writeOnlyChannel(ch, 160)
	v = readOnlyChannel(ch)
	fmt.Printf("v from read only channel %d\n", v)

	close(ch) // close can only happen on the provider side, not consumer side
	// consume the items in the channel, if the channel is empty but not closed,
	// the excution will be suspended until new items comming in, until deadlock
	// for loop will only exit when no items in the channel, and the channel is closed
	for ele := range ch {
		fmt.Printf("consuming %d\n", ele)
	}
}

// write only channel
func writeOnlyChannel(ch chan<- int, val int) {
	ch <- val
}

func readOnlyChannel(ch <-chan int) int {
	return <-ch
}

func ChannelRun() {
	rangeChannel()
}
