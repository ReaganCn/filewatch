package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func Exit(format string, usage string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, filepath.Base(os.Args[0])+": "+format+"\n", a...)
	fmt.Print("\n" + usage)
	os.Exit(1)
}
