package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

func main() {
	fmt.Println("Belajar Struct ")

	var s1 student

	s1.name = "Aries Dimas"
	s1.grade = 100

	fmt.Println(s1)

	var s2 = student{"Ethan Hunt", 80}

	fmt.Println(s2)

	var s3 = student{grade: 67, name: "Peter Griffin"}

	fmt.Println(s3)

}
