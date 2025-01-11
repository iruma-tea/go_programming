package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v", i, i)
}
