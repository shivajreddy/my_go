package main

import (
	"fmt"
)

// given a channel, send data to that channel
func ping(ch chan int) {
	ch <- 420
}

// given a channel, send data to that channel
func ping2(ch chan int) {
	ch <- 69
}

// get data from channel, modify it, send it back
func pong(ch chan int) {
	val := <-ch
	val += 1
	ch <- val
}

func pong2(ch chan int) {
	val := <-ch
	val += 20
	ch <- val
}

// print all the messges in the channel
func print_chan_messges(ch chan int) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func hub() {
	ch := make(chan int) // create a channel
	go print_chan_messges(ch)

	/*
		go func() {
			ch <- 420
		}()
		val := <-ch
		val += 1
		go func() {
			ch <- val
		}()
		// */

	// /*
	ch <- 420
	ch <- 421
	// */

	close(ch)
}
