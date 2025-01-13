package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v2 := Vertex{X: 1}
	fmt.Println(v2)
}
