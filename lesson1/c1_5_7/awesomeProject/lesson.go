package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("innner func", x)
	}(1)
}
