package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("fdafdsafa")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!!")

	fmt.Println("ok!")

}
