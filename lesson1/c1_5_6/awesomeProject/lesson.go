package main

import "fmt"

func main() {
	f := func(x int) {
		fmt.Println("innner func", x)
	}
	f(1)
}
