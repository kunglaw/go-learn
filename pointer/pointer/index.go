package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

func (s student) changeName(name string) {
	fmt.Println(" ---> on changeName1, name changed to", name)
	s.name = name
	fmt.Println("inside changeName1 method ", s.name)
}

func (s *student) changeName2(name string) {

	fmt.Println("---> on changeName2 , name changed to", name)
	s.name = name
}

func main() {
	var s1 = student{"Aries Dimas", 30}
	fmt.Println("s1 before", s1.name)

	s1.changeName("jason born")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeName2("Ethan Hunt")
	fmt.Println("s1 afetr changeName2", s1.name)
}
