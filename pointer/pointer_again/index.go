package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = a

	fmt.Println("b => ", b)
	fmt.Println("a => ", a)

	b = 12
	fmt.Println("b => ", b)
	fmt.Println("a => ", a)

	var c *int = &a /* initialisasi memori nilai a */
	*c = 56
	fmt.Println("c => ", c)
	fmt.Println("*c => ", *c)
	fmt.Println("a => ", a)
}
