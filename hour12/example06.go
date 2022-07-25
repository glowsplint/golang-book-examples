package main

import (
	"fmt"
	"time"
)

func ping(i int, c chan string) func(c chan string) {
	// Function factory
	return func(c chan string) {
		for {
			time.Sleep(1 * time.Second)
			c <- fmt.Sprintf("ping on channel %v", i)
		}
	}
}

func main() {
	// Using a select statement, with timeout
	channel1 := make(chan string)
	channel2 := make(chan string)

	ping1 := ping(1, channel1)
	ping2 := ping(2, channel1)

	go ping1(channel1)
	go ping2(channel2)

	select {
	case msg1 := <-channel1:
		fmt.Println("received", msg1)
	case msg2 := <-channel2:
		fmt.Println("received", msg2)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("no messages received, giving up...")
	}

}
