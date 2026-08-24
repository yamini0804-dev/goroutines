package main

import (
	"fmt"
	"sync"
)

var data = make(map[string]int)
var mutex sync.RWMutex

func readData(key string, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.RLock()
	fmt.Println("reading:", data[key])
	mutex.RUnlock()
}

func writeData(key string, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	data[key] = value
	mutex.Unlock()
}
func main() {
	var wg sync.WaitGroup

	data["A"] = 100

	wg.Add(4)

	go readData("A", &wg)
	go readData("A", &wg)
	go readData("A", &wg)

	go writeData("A", 200, &wg)

	wg.Wait()

	fmt.Println("Final:", data["A"])
}
