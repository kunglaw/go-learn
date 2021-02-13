package main

import "fmt"

// import "string"

func main() {

	var x string = ""

	if x == "dimas" {
		fmt.Println(" Yes this is Dimas")
	} else if x == "rian" {
		fmt.Println("Ini Rian")
	} else {
		fmt.Println("Bukan siapa siapa")
	}

	// temporary variable pada if

	var point = 1000.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect \n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s percent \n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad \n", percent, "%")
	}

}
