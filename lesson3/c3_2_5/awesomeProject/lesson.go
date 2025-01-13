package main

import "fmt"

func main() {
	var p2 *int
	fmt.Println(p2)
	*p2++
	fmt.Println(p2)
}
