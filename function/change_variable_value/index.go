package main

import (
	"fmt"
)

func main() {
	score := 10
	fmt.Print(score)

	change(100)
	fmt.Println(score)

}

func change(val int) {
	score = val
}
