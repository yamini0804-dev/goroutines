// sync.once initalize once when multiple go routines are called
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 5; i++ {

		go worker(&wg)
	}
	wg.Wait()

}
func initialize() {
	fmt.Println("initialization performed")
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initialize)
}
