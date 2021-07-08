package main

import (
	"fmt"
	"go-learn/exportImport/library2"
)

func main() {
	var s2 = library2.Student{"Aries Dimas", 30}
	fmt.Println("name", s2.Name)
	fmt.Println("grade", s2.Grade)
}
