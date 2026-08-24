package main

import (
	"fmt"
	"sync"
)

var counter int
var mut sync.Mutex

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
