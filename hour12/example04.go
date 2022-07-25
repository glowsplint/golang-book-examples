package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	// Sends "ping" onto the channel every second, regular polling
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	go pinger(messages)
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
