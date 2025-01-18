package main

import "fmt"

func do(i interface{}) {
	ss := i.(string)
	fmt.Println(ss + "!")
}

func main() {
	do("Mike")
}
