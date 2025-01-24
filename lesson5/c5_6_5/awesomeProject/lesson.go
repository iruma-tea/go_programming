package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(599 * time.Millisecond)

	select {
	case <-tick:
		fmt.Println("tick.")
	case <-boom:
		fmt.Println("BOOM!")
	default:
		fmt.Println("    .")
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("##############")
}
