package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name":"Aries Dimas","Age":32}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.Fullname)
	fmt.Println("age :", data.Age)

}
