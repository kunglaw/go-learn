package main

import (
	"fmt"
	"reflect"
)

func main() {

	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel : ", reflectValue.Type())
	fmt.Println(reflectValue.Kind())
	fmt.Println("nilai variabel:", reflectValue.Int())

}
