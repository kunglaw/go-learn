package main

import (
	"fmt"
	"time"
)

func main() {
	var number int = 1_000_000
	start := time.Now()
	fmt.Println(" start => ", start)

	for i := 0; i < number; i++ {
		fmt.Println(i)
	}

	elapsed := time.Since(start)
	fmt.Println(" end => ", elapsed)
}
