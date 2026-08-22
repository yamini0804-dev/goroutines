package main

import ( // see  about it
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Worker started")

	time.Sleep(2 * time.Second)

	fmt.Println("Worker finished")

	done <- true
}

func main() {
	done := make(chan bool)

	go worker(done)

	<-done

	fmt.Println("Main finished")
}
