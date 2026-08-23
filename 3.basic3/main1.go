package main

import (
	"fmt"
	"time"
)

func main() {
	go printnums()
	fmt.Println("main function starts")
	time.Sleep(8 * time.Millisecond)
	fmt.Println("main function finishes")
}

func printnums() {
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}
	time.Sleep(4 * time.Millisecond)
}
