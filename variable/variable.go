package main

import "fmt"

func main() {
	var firstName = "Aries Dimas"
	var lastName = "Yudhistira"

	fmt.Println(firstName + " " + lastName)
	fmt.Printf("Halo %s %s \n", firstName, lastName)

	// deklarasi variable tanpa tipe data

	var statement string = "lorem ipsum sit dolor amet" // declare data type
	statementAgain := "iam give up"

	fmt.Printf(" what he say ? %s %s! \n", statement, statementAgain)

	// deklarasi multi variabel

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	fmt.Println(first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	fmt.Println(fourth, fifth, sixth)

	// underscore untuk variable yang tidak terpakai 

	_ = 123
	_ = "hello unused variable"

}
