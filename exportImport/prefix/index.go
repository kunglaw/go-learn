package main

import (
	"fmt"
	. "go-learn/exportImport/library2"
)

func main() {
	var s1 = Student{"Aries Dimas", 30}
	fmt.Println("name", s1.Name)
	fmt.Println("grade", s1.Grade)

}
