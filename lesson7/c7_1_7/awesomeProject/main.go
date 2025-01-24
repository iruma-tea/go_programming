package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	// fmt.Println(req)
}
