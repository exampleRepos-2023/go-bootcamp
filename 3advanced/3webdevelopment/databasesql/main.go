package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Database SQL
	// Package sql

	// Installing PostgreSQL (or SQLite, MySQL...)
	// https://www.postgresql.org/

	// Pull the database package (after go mod init)
	// go get .github.com/lib/pq

	// Import to connect package sql to PostgreSQL

	// Open database connection
	const (
		host     = "localhost" // Default
		port     = 5432        // Default
		user     = "postgres"  // Default
		password = "postgres"  // Custom
		dbname   = "postgres"  // Default
	)
	connectionInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionInfo)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)

	createTable(db)
	runInsert(db)
	runSelect(db)
	dropTable(db)
}

func createTable(db *sql.DB) {
	createS := "CREATE TABLE Weapons(Name varchar(255), Level int);"
	_, err := db.Exec(createS)
	checkError(err)
}

func runInsert(db *sql.DB) {
	insertS := "INSERT INTO Weapons(Name, Level) VALUES('Ninja Star', 1);"
	_, err := db.Exec(insertS)
	checkError(err)
}

func runSelect(db *sql.DB) {
	selectS := "SELECT Name, Level FROM Weapons;"
	rows, err := db.Query(selectS)
	checkError(err)
	for rows.Next() {
		var name string
		var level int
		err := rows.Scan(&name, &level)
		fmt.Println(name, level, err)
	}
}

func dropTable(db *sql.DB) {
	dropS := "DROP TABLE IF EXISTS Weapons;"
	_, err := db.Exec(dropS)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
