package main

import "fmt"

func do(i interface{}) {
	// i *= 2 // 型アサーションをしないとエラーとなる。
	// fmt.Println(i)
	ii := i.(int)
	ii *= 2
	fmt.Println(ii)
}

func main() {
	var i interface{} = 10
	do(i)
}
