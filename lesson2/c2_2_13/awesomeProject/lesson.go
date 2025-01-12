package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}

	for _, v := range m {
		fmt.Println(v)
	}
}
