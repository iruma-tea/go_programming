package main

import "fmt"

func getOsName() string {
	return "mac"
}

func main() {
	os := getOsName()
	switch os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
		// default:
		// 	fmt.Println("Default!!")
	}
}
