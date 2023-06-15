package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("SELECT id, name, grade from tb_student WHERE age > ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.id, each.name)
	}
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	stmt, err := db.Prepare("SELECT name, grade from tb_student WHERE id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name :%s \n grade :%d \n ", result1.name, result1.grade)

	var result2 = student{}
	stmt.QueryRow("B001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name :%s \n grade :%d \n ", result2.name, result2.grade)
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}

	var id string = "E001"

	err = db.QueryRow("SELECT name, grade FROM tb_student WHERE id = ? ", id).
		Scan(&result.name, &result.grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s \n grade: %d \n", result.name, result.grade)
}

func main() {
	// sqlQuery()
	// sqlQueryRow()
	sqlPrepare()
}
