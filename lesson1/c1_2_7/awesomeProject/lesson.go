package main

import "fmt"

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 2
	xf64 := 1.3
	xs := "test test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Println(i, f64, s, t, f)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
