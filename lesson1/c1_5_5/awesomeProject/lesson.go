package main

import "fmt"

func cal(price, item int) (result int) {
	result = price * item
	return result
}

func main() {
	r := cal(100, 2)
	fmt.Println(r)
}
