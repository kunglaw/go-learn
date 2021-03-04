package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
	age  int
}

type employee struct {
	salary   int
	position string
	person
}

func main() {

	var p1 = person{
		name: "Dimas",
		age:  30,
	}

	var s1 = employee{
		salary:   12000000,
		position: "Web Developer",
		person:   p1,
	}

	// anonymous struct
	var s2 = struct {
		person
		pangkat string
	}{
		pangkat: "Sersan Dua",
		person: person{
			name: "Pramuhin",
			age:  50,
		},
	}

	fmt.Println(s1)

	fmt.Println("person name : ", s1.person.name)
	fmt.Println("person age : ", s1.person.age)

	fmt.Println(s2)

	print(json.Marshal(s2))

}
