package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("say hi")

	db := createAndOpen("testdb")
	defer db.Close()
	
	err := db.Ping()
	if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
	}	
	
	rows, err := db.Query("SELECT id, name FROM test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rows, err)

	result, err := db.Exec("INSERT INTO test (id, name) VALUES (4, 'Alex4')")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Insert result: ", result)
}

func createAndOpen(dbName string) *sql.DB {
	db, err := sql.Open("mysql", "root:2831647d@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql",  "root:2831647d@tcp(127.0.0.1:3306)/" + dbName)
	if err != nil {
		panic(err)
	}
	return db
}