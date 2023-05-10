package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	models "github.com/reagancn/filewatch/pkg/database/models"
)

func RunCLI(db *sql.DB) {

	var path string

	// Define greet command line flags
	setfs := flag.NewFlagSet("set", flag.ExitOnError)
	setfs.StringVar(&path, "path", "", "The path of the file to watch")

	// Check for command argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		os.Exit(1)
	}

	// Execute command based on first argument
	switch os.Args[1] {
	case "set":
		setfs.Parse(os.Args[2:])
		set(path, db)
		break

	case "run":
		run()
		break

	default:
		fmt.Println("Invalid command specified. Usage is: ", usage)
		os.Exit(1)
	}
}

func set(path string, db *sql.DB) {

	// Initialize the path model
	p := models.Path{Path: path}

	// Check if required flags are set
	if path == "" {
		fmt.Println("Please provide a path to the file to watch")
		os.Exit(1)
	}

	// Insert the path into the database
	p.Insert(db)
}

func run() {
	log.Println("Watching files...")

	// TODO: Execute the file watcher here
}
