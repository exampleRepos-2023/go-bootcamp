package persistence

import (
	"database/sql"
	"fmt"
	_ ".github.com/lib/pq"
	"log"
)

const (
	host     = "localhost" // Default
	port     = 5432        // Default
	user     = "postgres"  // Default
	password = "postgres"  // Custom
	dbname   = "postgres"  // Default
)

var connectionInfo = fmt.Sprintf(
"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)

var DB = newDatabase()

func newDatabase() *sql.DB {
	db, err := connectDatabase()
	if err != nil {
		log.Println("Unable to connect database", err)
		return nil
	}

	dropTable(db, questionsTable)
	dropTable(db, answersTable)
	dropTable(db, questionAnswersTable)

	creatTable(db, questionsTable, createQuestionsTable)
	creatTable(db, answersTable, createAnswersTable)
	creatTable(db, questionAnswersTable, createQuestionAnswersTable)

	return db
}

func connectDatabase() (*sql.DB, error) {
	connectionInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Println("Unable to open database connection", err)
		return nil, err
	}
	return db, nil
}

func creatTable(db *sql.DB, tableName string, createQuery string) {
	_, err := db.Exec(createQuery)
	if err != nil {
		log.Println("Unable to create database table", err)
		dropTable(db, tableName)
	}
}

func dropTable(db *sql.DB, tableName string) {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Unable to drop database table", err)
	}
}

func closeDatabase(db *sql.DB) {
	dropTable(db, questionsTable)
	dropTable(db, answersTable)
	dropTable(db, questionAnswersTable)

	err := db.Close()
	if err != nil {
		log.Println("Unable to close database connection", err)
	}
}