package main

import (
	"fmt"
	"math"
)

func main() {
	var area, circumference = calculate(10)
	// test.test()
	fmt.Println(area, circumference)
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)

	var circumference = math.Pi * d

	return area, circumference
}
