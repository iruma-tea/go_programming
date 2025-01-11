package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
}
