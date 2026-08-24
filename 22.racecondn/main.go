package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("counter:", counter)

}
