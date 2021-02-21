package main

import "fmt"

func main() {
	fmt.Println("data type")

	var positiveNumber uint8 = 89
	var negativeNumber = -34

	var decimalNumber = 2.62

	fmt.Printf(" Bilangan positif: %d \n ", positiveNumber)
	fmt.Printf(" Bilangan negative: %d \n ", negativeNumber)

	fmt.Println("Bilangan Desimal ", decimalNumber)

	// string

	var name string = "Dimas"
	var nameIndex string = string(name[2])

	fmt.Println(" huruf : ", name, "nameIndex : ", nameIndex)

}
