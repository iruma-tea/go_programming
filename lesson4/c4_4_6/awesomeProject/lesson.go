package main

func do(i interface{}) {
	// i.(type) // use of .(type) outside type switch
}

func main() {
	do("Mike")
}
