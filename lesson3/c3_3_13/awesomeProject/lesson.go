package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)
}
