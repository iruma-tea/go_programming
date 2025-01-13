package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)
}
