package main

import "fmt"

func main() {
	var p *int = new(int)
	fmt.Println(p)

	var p2 *int
	fmt.Println(p2)
}
