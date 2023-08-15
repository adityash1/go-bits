package main

import (
	"fmt"
)

func main() {
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	for {
		msg, ok := <-msgch
		if !ok {
			break
		}
		fmt.Println("the message -> ", msg)
	}

	for msg := range msgch {
		fmt.Println(msg)
	}
}
