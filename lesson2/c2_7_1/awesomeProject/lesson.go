package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to connect to DB")
}

func save() {
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}
