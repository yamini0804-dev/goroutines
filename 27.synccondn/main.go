package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var cond = sync.NewCond(&mutex)
var mutex sync.Mutex
var ready bool

func main() {
	go worker()
	time.Sleep(2 * time.Second)
	mutex.Lock()
	ready = true
	fmt.Println("main:data is ready")
	cond.Signal()
	mutex.Unlock()
	time.Sleep(1 * time.Second)

}

func worker() {
	mutex.Lock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("worker data is ready")
	mutex.Unlock()

}
