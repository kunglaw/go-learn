package main

import "fmt"

func main() {
	// fmt.Print("Bisa gak tanpa main ?")

	person := map[string]string{
		"name":  "Dimas",
		"age":   "30",
		"skill": "programming",
	}

	//any type

	personAgain := map[string]interface{}{
		"name":  "Aries Dimas",
		"age":   30,
		"skill": "Web Development",
	}

	// array of person
	var people = [...]string{"dimas", "farida", "baskoro"}

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

	fmt.Println(person, personAgain)
	fmt.Println(people)
	fmt.Println(advancePeople)

	// looping array
	for i := 0; i < len(people); i++ {
		fmt.Println(people[i])
	}

	for j := 0; j < len(advancePeople); j++ {
		fmt.Println(advancePeople[j])
	}

	detectionItem()

}

func detectionItem() {
	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	fmt.Println(value, isExist) // kenapa value 0 ?

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}

func sliceMapCombination() {

}
