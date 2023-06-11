package main

import "fmt"

func welcome() {
	fmt.Println("Kono Dio daa")
	defer fmt.Println("halo")
	fmt.Println("Selamat Datang")
}

func main() {
	orderSomeFood("pizza")
	orderSomeFood("burger")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terima Kasih, silakan tunggu")

	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak", "\n")
		return // return artinya berhenti dari blok code
	}

	fmt.Println("Pesanan anda: ", menu)
}
