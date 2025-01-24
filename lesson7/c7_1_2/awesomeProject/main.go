package main

import (
	"fmt"
	"net/url"
)

func main() {
	base, _ := url.Parse("http://example.com")
	fmt.Println(base)
}
