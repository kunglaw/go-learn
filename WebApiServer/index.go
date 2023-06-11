package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var students = []student{
	student{"001", "Aries Dimas", 100},
	student{"002", "Ethan Hunt", 90},
	student{"003", "Donal Trump", 60},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(students)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var getQuery = r.URL.Query()
		var getParams = r.URL.Path

		var id = getQuery.Get("id")
		//fmt.Println(id)
		fmt.Println(getParams, id)
		var result []byte
		var err error

		for _, each := range students {
			fmt.Println(each, id)
			if each.ID == id {

				result, err = json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {

	http.HandleFunc("/user", user)
	http.HandleFunc("/users", users)

	fmt.Println("starting web server at http://localhost:8182/")

	http.ListenAndServe(":8182", nil)

}
