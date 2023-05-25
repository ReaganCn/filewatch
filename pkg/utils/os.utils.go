package utils

import (
	"fmt"
	"os"
)

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dir
}
