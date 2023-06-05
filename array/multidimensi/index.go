package main

import (
	"encoding/json"
	"fmt"
)

func multiGenerate(num int) {

	var multi = [][]int{}

	init := 1
	for i := 0; i < num; i++ {
		var temp []int
		for j := 0; j < num; j++ {
			temp = append(temp, init)
			// temp[j] = init
			init++
		}
		multi = append(multi, temp)
	}

	//fmt.Printf("%+q", multi)

	b, err := json.MarshalIndent(multi, "", "  ")

	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err)
	}

}

func main() {
	fmt.Println("multidimensi")

	multiGenerate(5)
}
