package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./filewatch.db")

	if err != nil {
		log.Println("Error opening database")
		log.Fatal(err)
	}

	defer db.Close()

	// Enable foreign key support
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Println("Error enabling foreign key support")
		log.Fatal(err)
	}

	return db, err
}

func Get(db *sql.DB, query string) (*sql.Rows, error) {

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error querying database")
		log.Fatal(err)
	}
	return rows, err
}

func Insert(db *sql.DB, tablename string, columns []string, values []string) (sql.Result, error) {

	// Convert the columns and values to strings
	strColumns := strings.Join(columns, ", ")
	strValues := strings.Join(values, ", ")

	// Create the insert statement
	query := "INSERT INTO " + tablename + " (" + strColumns + ") VALUES (" + strValues + ")"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error preparing insert statement")
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec()
	if err != nil {
		log.Println("Error executing insert statement")
		log.Fatal(err)
	}

	// Get the auto-generated ID of the inserted row
	insertedID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inserted data with ID: %d\n", insertedID)

	return result, err
}
