package main

import "fmt"

type MyInt int

func (i MyInt) Dobule() int {
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Dobule())
}
