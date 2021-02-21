package main

import "fmt"

func main() {

	var getMinMax = func(n []int) (int, int) {
		var min, max int = 0, 0

		for i, e := range n {

			//fmt.Println("e ", e)
			if i == 0 {
				max, min = e, e
			} else if e > max {
				max = e
			} else if e < min {
				min = e
			}
		}

		return max, min

	}

	var data = []int{3, 9, 3, 4, 6, 2, 5, 8, 9, 7, 4}

	var max, min = getMinMax(data)

	fmt.Println(data)
	fmt.Println("max => ", max, "min => ", min)

}
