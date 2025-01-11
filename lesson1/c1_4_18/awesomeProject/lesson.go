package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	m["banana"] = 300
	fmt.Println(m)
}
