package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Hello World"
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))
}
