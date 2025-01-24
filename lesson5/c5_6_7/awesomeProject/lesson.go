package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		isBreak := false
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			isBreak = true
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
		if isBreak {
			break
		}
	}
	fmt.Println("##############")
}
