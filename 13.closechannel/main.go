package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 5; i++ {
			ch <- i
		}

		close(ch)
	}()

	for {
		value, ok := <-ch
		if !ok { // channel will have value and ok  if value = 0 ok false
			fmt.Println("channel closed")
			break
		}
		fmt.Println(value)
	}
}
