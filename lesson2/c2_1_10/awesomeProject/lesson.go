package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}
	fmt.Println(result)

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great2")
	}
	// fmt.Println(result2) // エラー
}
