package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalln("error!!")

	fmt.Println("ok!")
}
