package main

import "fmt"

func main() {

	var myList = []map[string]interface{}{
		{
			"name": "Aries Dimas",
			"age":  30,
		},
		{
			"name":   "Asti",
			"age":    30,
			"salary": 12_000_000,
		},
		{
			"name":    "Fadil",
			"company": "PT Giyanta",
		},
	}

	var typedList = []map[string]interface{}{
		{
			"name": "Aries Dimas",
			"age":  30,
		}, {
			"name": "Asti",
			"age":  31,
		},
	}

	fmt.Println(myList)
	fmt.Println(typedList)

}
