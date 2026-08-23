// we have 6 workers they are going to perform three tasks
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(6)
	for i := 1; i <= 6; i++ {
		go workers(i, &wg)

	}
	wg.Wait()
	fmt.Println("finished")
}

func workers(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := 1; task <= 4; task++ {
		fmt.Println("workers", id, "performing", task)
	}
	time.Sleep(500 * time.Millisecond)
}
