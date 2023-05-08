package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func FilewatchCLI() {

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
		set(path)
		break

	case "run":
		run()
		break

	default:
		fmt.Println("Invalid command specified. Usage is: ", usage)
		os.Exit(1)
	}
}

func set(path string) {
	// Check if required flags are set
	if path == "" {
		fmt.Println("Please provide a path to the file to watch")
		os.Exit(1)
	}

	// TODO: Set the path to an sqlite database here
}

func run() {
	log.Println("Watching files...")

	// TODO: Execute the file watcher here
}
