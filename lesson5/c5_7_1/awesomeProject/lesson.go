package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(map[string]int)
	go func() {
		for i := 0; i < 10; i++ {
			c["key"] += 1
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c["key"] += 1
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c["key"])
}
