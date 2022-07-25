package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	stop := make(chan bool)
	go sender(messages)
	go func() {
		// Functionally the same as reading from the time.After channel
		// However, now any goroutine can write to this stop channel which is useful
		// This allows us to terminate the for loop via early return from any goroutine
		time.Sleep(time.Second * 3)
		fmt.Println("Time's up!")
		stop <- true
	}()

	for {
		select {
		case <-stop:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
