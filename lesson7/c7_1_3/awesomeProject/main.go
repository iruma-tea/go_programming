package main

import (
	"fmt"
	"net/url"
)

func main() {
	base, err := url.Parse("http://e xample.com")
	fmt.Println(base, err)
}
