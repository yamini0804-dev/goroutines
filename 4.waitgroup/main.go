package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go prinstmsg(&wg)
	go printnums(&wg)
	wg.Wait()

	fmt.Println("all the goroutines are finished")
}

func prinstmsg(wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Println(string(ch))
		time.Sleep(200 * time.Millisecond)
	}
}

func printnums(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}
