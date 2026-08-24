package main

// see more not understood
import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("worker", id, "processing", job)

		results <- job * 2

	}

}
