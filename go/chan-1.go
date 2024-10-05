package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello via channel"
	}()
	fmt.Println(<-stringStream)
}
