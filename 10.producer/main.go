package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

}

func consumer(ch chan int) {
	for values := range ch {
		fmt.Println("recieved", values)

	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
