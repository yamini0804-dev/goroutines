//  5 workers and they are performing three tasks

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i <= 3; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("function is finished")

}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := 0; task <= 5; task++ {
		fmt.Println("works", id, "processing task", task)
		time.Sleep(500 * time.Millisecond)
	}

}
