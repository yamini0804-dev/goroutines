package main

import (
	"fmt"
	"time"
)

func printMessage(id int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Goroutine", id, ":", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage(1)
	go printMessage(2)
	go printMessage(3)
	go printMessage(4)
	go printMessage(5)

	time.Sleep(3 * time.Second)
}
