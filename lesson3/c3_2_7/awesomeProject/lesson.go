package main

import "fmt"

func main() {
	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
