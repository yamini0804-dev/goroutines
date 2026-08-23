package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Millisecond)
		ch <- "finished"

	}()
	select {
	case msg := <-ch:
		fmt.Println(msg)

	case <-time.After(2 * time.Second):
		fmt.Println("timed out") // channel that is gping to become ready after two seconds
	}
}
