package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}
