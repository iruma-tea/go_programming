package main

import (
	"fmt"
	"time"
)

func goroutine(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go goroutine("world")
	normal("hello")
	time.Sleep(2000 * time.Millisecond)
}
