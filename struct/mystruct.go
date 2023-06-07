package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

type product struct {
	name        string
	price       int
	description string
}

func main() {
	fmt.Println("Belajar Struct ")

	var s1 student
	var s4 student
	var p1 product

	s1.name = "Aries Dimas"
	s1.grade = 100

	s4.name = "Rian D"
	s4.grade = 80

	p1.name = "GTX 1660 6GB"
	p1.price = 8000000
	p1.description = "this for game only not for mining"

	fmt.Println(s1)

	var s2 = student{"Ethan Hunt", 80}

	fmt.Println(s4, s2)

	var s3 = student{grade: 67, name: "Peter Griffin"}

	fmt.Println(s3)

	fmt.Print(p1.description)

}
