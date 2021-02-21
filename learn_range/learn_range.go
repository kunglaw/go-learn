package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 4, 5, 5, 3, 2, 5, 3, 2, 5, 6, 1}

	for i, c := range a {
		fmt.Println(c, i)
	}

	var advancePeople = [...]map[string]interface{}{
		{
			"name":  "Aries Dimas",
			"age":   30,
			"skill": "Web Development",
		},
		{
			"name":  "Asti Sudibyo",
			"age":   31,
			"skill": "ERP Consultant",
		},
		{
			"name":  "Baskoro",
			"age":   31,
			"skill": "Electric Engineering",
		},
		{
			"name":  "Farida",
			"age":   31,
			"skill": "Data Scientist",
		},
	}

	for k, v := range advancePeople {
		fmt.Println(k, v)
		for k1, v1 := range v {
			fmt.Println(k1, ":", v1)
		}
	}
}
