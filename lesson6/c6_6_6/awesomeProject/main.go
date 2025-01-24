package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a([a-z0-9]+)e", "appl0e")
	fmt.Println(match)
}
