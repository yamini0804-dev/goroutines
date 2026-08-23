// send only channel
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(1, ch, &wg)
	go producer(2, ch, &wg)
	go func() {
		wg.Wait()
		close(ch)
	}()
	for value := range ch {
		fmt.Println("recived:", value)
	}

}

func producer(id int, ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 3; i++ {
		ch <- id*10 + i
	}
}
