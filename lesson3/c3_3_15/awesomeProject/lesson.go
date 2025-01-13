package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	(*v).X = 1000
}

func main() {
	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)
}
