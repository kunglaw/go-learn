package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var names = []string{"Aries", "Dimas"}
	// var people = ["aries","dimas"]
	// fmt.Println(names)
	printMessage("halo ", names)
	// library.functionReturn()
	// functionReturn()
	log.Println("Ini Print dari log") // 2021/02/16 22:08:07 Ini Print dari log
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
