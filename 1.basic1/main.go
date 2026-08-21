package main

import (
	"fmt"
	"time"
)

func main() {
	go printnumbers()
	go printchars()
	time.Sleep(3 * time.Second)
}

func printnumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(3 * time.Millisecond)
	}

}

func printchars() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Println(string(ch))
		time.Sleep(3 * time.Millisecond)

	}
}
