package main

import "time"

func main() {
	// go printnums()
	go printchar()
	go printnums()
	time.Sleep(3 * time.Second)

}

func printnums() {
	for i := 0; i <= 3; i++ {
		print(i)
	}
	time.Sleep(3 * time.Second)
}

func printchar() {
	for ch := 'a'; ch <= 'e'; ch++ {
		print(string(ch))
	}
	time.Sleep(2 * time.Millisecond)
}
