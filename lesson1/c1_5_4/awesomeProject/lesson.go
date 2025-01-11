package main

import "fmt"

func add(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)
}
