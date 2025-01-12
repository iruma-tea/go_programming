package main

import "fmt"

func getOsName() string {
	return "mac"
}

func main() {
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!", os)
	}
	// fmt.Println(os) // Error: undefined: os
}
