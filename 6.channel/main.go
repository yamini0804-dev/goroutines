package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello" // pour into channel
	}()

	message := <-ch // drink out of channel
	fmt.Println(message)

}
