package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Halo", s.name)
}

func (s student) getName(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var s1 = student{"Aries Dimas", 30}
	s1.sayHello()
	var myName string = s1.getName(2)
	fmt.Println(myName)
}
