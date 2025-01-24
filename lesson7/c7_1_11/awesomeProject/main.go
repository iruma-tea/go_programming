package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
