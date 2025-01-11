package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "Hello World"
	fmt.Println(s)
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
}
