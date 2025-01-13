package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func main() {
	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)
}
