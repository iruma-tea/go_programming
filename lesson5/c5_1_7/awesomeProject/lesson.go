package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)
	normal("hello")
	wg.Wait()
}
