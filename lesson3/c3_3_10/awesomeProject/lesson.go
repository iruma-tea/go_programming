package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func main() {
	v6 := new(Vertex)
	fmt.Println(v6)
}
