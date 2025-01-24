package main

import (
	"fmt"
	"regexp"
)

func main() {
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("e/view/test")
	fmt.Println(fs)
}
