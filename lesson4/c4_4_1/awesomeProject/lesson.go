package main

func do(i interface{}) {
	// 以下、エラーとなる。
	// ii := i * 2
	// fmt.Println(ii)
}

func main() {
	do(10)
}
