package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex
var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mut.Lock()
		counter++
		mut.Unlock()

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
