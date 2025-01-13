package main

import "fmt"

func main() {
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
}
