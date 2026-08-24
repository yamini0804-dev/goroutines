package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://youtube.com",
		"https://github.com",
	}
	var wg sync.WaitGroup

	wg.Add(len(urls)) // jitne url utne go routines pass honge

	for _, url := range urls {
		go fetch(url, &wg)
	}
	wg.Wait()
	fmt.Println("requests done")
}

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer response.Body.Close()
	fmt.Println(response.Status)
}
