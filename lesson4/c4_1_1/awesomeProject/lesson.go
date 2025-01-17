package main

import "fmt"

type Vertex struct {
	X, Y int
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
}
