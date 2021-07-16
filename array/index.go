package main

import "fmt"

func createMultiDimension() {

	// var side int = 5
	// var results = [5][5]int{}
	// var strings = [5][5]string{}
	var filledArr = []string{}

	fmt.Println(filledArr)
	filledArr = append(filledArr, "Hello")

	// for i := 0; i < side; i++ {
	// 	results[i] = [5]int{}
	// 	for j := 0; j < side; j++ {
	// 		results[i][j] = i
	// 	}

	// }

	fmt.Println(filledArr)

}

func multiDimension() {
	var numbers = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	for i := 0; i <= len(numbers[0])-1; i++ {
		for j := 0; j < len(numbers[0])-1; j++ {
			numbers[i][j] = numbers[i][j] * 2
		}
	}

	fmt.Println(numbers)
}

func main() {
	// multiDimension()
	createMultiDimension()
}
