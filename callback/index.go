package main

import "fmt"

func foo(a string, age int, experiences int, cb func(age int, experiences int)) {
	fmt.Println("Hello " + a)
	cb(age, experiences)
}

func bar(age int, experiences int) {
	fmt.Println("age => ", age)
	fmt.Println("experiences => ", experiences)
}

func main() {
	foo("Aries Dimas", 32, 10, bar)
}
