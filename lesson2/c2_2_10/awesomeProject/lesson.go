package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for _, v := range l {
		fmt.Println(v)
	}
}
