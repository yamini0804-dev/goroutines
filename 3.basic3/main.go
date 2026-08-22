package main

import (
	"fmt"
	"time"
)

func main() {
	go printmsg()
	fmt.Println("main function started")
	time.Sleep(2 * time.Second)
	fmt.Println("main function finished")

}

func printmsg() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

// as the main function stops the go routines laso stops
