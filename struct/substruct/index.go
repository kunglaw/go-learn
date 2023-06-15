package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	grade int
}

func main() {
	var p1 = person{name: "Dimas", age: 30}
	var p2 = person{name: "Kunglaw", age: 32}
	p2.name = "Babe"
	fmt.Println(p2)
	var s1 = student{person: p1, grade: 2}

	fmt.Println(s1)

}
