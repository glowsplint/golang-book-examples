package main

import (
	"log"
)

func main() {
	// Logging to a file directly via the operating system
	// go run example04.go > example04.log
	for i := 1; i <= 5; i++ {
		log.Printf("Log iteration %d", i)
	}
}
