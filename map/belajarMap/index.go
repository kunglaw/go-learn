package main

import (
	"fmt"
)

type myType interface {
	getName() string
	getNum() int
}

func main() {

	orang := map[int]interface{}{
		1: "Satu",
		2: "Dua",
		3: "Tiga",
	}

	fmt.Println(orang)
	fmt.Println(orang[1])

}
