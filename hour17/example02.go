package main

import (
	"flag"
	"fmt"
)

func main() {
	// Using the flag package to add command line arguments, default values and help text
	s := flag.String("s", "Hello world", "String help text")
	flag.Parse()
	fmt.Println("value of s:", *s)
}
