package main

import (
	"log"

	sqlite "github.com/reagancn/filewatch/pkg/database"
)

func main() {
	// Initialize the database
	db, err := sqlite.ConnectDB()

	if err != nil {
		log.Println("Error connecting to database")
		log.Fatal(err)
	}

	// Initialize the CLI
	RunCLI(db)
}
