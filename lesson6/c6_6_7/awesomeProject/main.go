package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)
}
