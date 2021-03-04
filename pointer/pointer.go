// package fmt
package main

import "fmt"

func pointer() {
	// var number *int
	// var name *string

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) : ", numberA)
	fmt.Println("numberA (address) : ", &numberA)

	fmt.Println("numberB (value) : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)

}

func main() {
	pointer()
}
