package main

import "fmt"

type Human interface {
	Say()
}

type Person struct {
	Name string
}

func (p Person) Say() {
	fmt.Println(p.Name)
}

func main() {
	var mike Human = Person{"Mike"}
	mike.Say()
}
