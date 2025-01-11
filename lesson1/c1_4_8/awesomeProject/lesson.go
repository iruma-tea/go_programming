package main

import "fmt"

func main() {
	var bord = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(bord)
	fmt.Println(bord[1])
	fmt.Println(bord[1][2])
}
