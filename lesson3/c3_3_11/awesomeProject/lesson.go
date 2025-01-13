package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func main() {
	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)
}
