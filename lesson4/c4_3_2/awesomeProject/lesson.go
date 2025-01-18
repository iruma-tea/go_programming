package main

import "fmt"

type Human interface {
	Say()
}

type Person struct {
	Name string
}

// 以下のメソッドをコメントにするとコンパイルエラーとなる
func (p Person) Say() {
	fmt.Println(p.Name)
}

func main() {
	var mike Human = Person{"Mike"}
	mike.Say()
}
