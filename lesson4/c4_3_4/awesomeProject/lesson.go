package main

import "fmt"

type Human interface {
	Say()
}

type Person struct {
	Name string
}

func (p *Person) Say() {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
}

func main() {
	var mike Human = &Person{"Mike"}
	mike.Say()
}
