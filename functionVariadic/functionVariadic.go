package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f ", avg)

	fmt.Println(msg)

	people("Aries", "Dimas", "Yudhistira")
}

// car numbers means can inserted into many integer
func calculate(numbers ...int) float64 {
	result := 0

	for _, number := range numbers {
		result += number
	}

	var avg = float64(result) / float64(len(numbers))

	return avg
}

func people(names ...string) {

	result := ""

	for _, name := range names {
		result += "and " + name
	}

	fmt.Print(result)
}
