package main

// capacity of 3

import "fmt"

func main() {
	ch := make(chan int, 3)
	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
