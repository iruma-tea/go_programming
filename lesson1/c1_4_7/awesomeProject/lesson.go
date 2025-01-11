package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	n = append(n, 100)
	fmt.Println(n)
	n = append(n, 200, 300, 400)
	fmt.Println(n)
}
