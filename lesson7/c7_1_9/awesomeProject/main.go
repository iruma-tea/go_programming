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
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()
}
