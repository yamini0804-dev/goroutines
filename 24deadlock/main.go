package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Println("sending")
	ch <- 10
	fmt.Println("sent")
}

// no reciver so deadlock
