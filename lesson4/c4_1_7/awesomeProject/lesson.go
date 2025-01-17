package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	// return i * 2 // castしないとエラーになる
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
