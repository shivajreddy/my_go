package main

import (
	"fmt"
)

func main() {
	fmt.Println("Chat App Started")

	// channels

	// unbuffered channel ::
	// 		1. doesnt have to specify the size
	//		3. has no storage
	// 		a send blocks until another goroutine is ready to receive
	// 		2. channel is blocked until it receives a message
	// buffered channel ::
	// 		1. channel IS NOT blocked unti

	// channels
	ch := make(chan int)          // unbuffered (these are defautl)
	ch2 := make(chan string, 100) // buffered (like elixir mailbox)

	// sending (blocks until received if unbuffered)
	ch <- 420

	// receiving (blocks until something is sent)
	val := <-ch2

	close(ch)
	close(ch2)

	hub()
}
