package main

import "fmt"

func do(i interface{}) {
	ii := i.(int)
	ii *= 2
	fmt.Println(ii)
}

func main() {
	do(10)
}
