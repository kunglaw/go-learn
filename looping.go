package main

import "fmt"

func main() {

	//kondisi di luar
	var x = 0

	for x < 5 {
		fmt.Println("angka ", x)
		x++
	}

	// kondisi lengkap
	for i := 0; i <= 10; i++ {
		fmt.Println("iteration : ", i)
	}

	// for tanpa argumen
	var y = 0

	for {
		fmt.Println("Angka ", y)

		y++
		if y == 5 {
			break
		}
	}

	// perulangan bersarang

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// pemanfaatan label dalam perulangan

	for i := 0; i < 5; i++ {
	outerLoop:
		for j := 0; j < 5; j++ {
			if j == 2 {
				break outerLoop
			}
			fmt.Print(" matriks [", i, "][", j, "]", "\n")
		}
	}

}
