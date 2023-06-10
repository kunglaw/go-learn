package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(10)

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}
