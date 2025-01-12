package main

import "fmt"

func getOsName() string {
	return "dafdafad"
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
}
