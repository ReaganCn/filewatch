package main

import "fmt"

// import (
// 	"log"

// 	sqlite "github.com/reagancn/filewatch/pkg/database"
// )

func main() {
	// // Initialize the database
	// db, err := sqlite.ConnectDB()

	// if err != nil {
	// 	log.Println("Error connecting to database")
	// 	log.Fatal(err)
	// }

	// // Initialize the CLI
	// RunCLI(db)

	fmt.Printf("Welcome to %s 🥳! With filewatch you can monitor multiple files and directories for changes.\n", titleStyle.Render("filewatch"))
	fmt.Print("You can also set a custom webhook that will be triggered in the case of any event\n\n")
	bubble()
}
