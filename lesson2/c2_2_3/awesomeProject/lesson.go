package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
}
