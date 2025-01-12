package main

import "fmt"

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
	defer fmt.Println("world")
	foo()
	fmt.Println("hello")
}
