package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to connect to DB")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}
