package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
