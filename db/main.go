package main

import (
	"database/sql"
	"fmt"
	"os"

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

	result, err := db.Exec("INSERT INTO test (id, name) VALUES (5, 'Alex5')")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Insert result: ", result)
}

func createAndOpen(dbName string) *sql.DB {
	host := os.Getenv("RDS_HOSTNAME")
	userName := os.Getenv("RDS_USERNAME")
	password := os.Getenv("RDS_PASSWORD")
	port := os.Getenv("RDS_PORT")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/", userName, password, host, port)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql",  dataSourceName + dbName)
	if err != nil {
		panic(err)
	}
	return db
}