package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I dont't know about type %T!\n", v)
	} // i.(type) // use of .(type) outside type switch
}

func main() {
	do(10)
	do("Mike")
	do(true)
}
