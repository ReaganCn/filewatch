package main

import (
	"log"
	"os"

	"github.com/reagancn/filewatch/pkg/utils"
)

func StepsHandler(m model) (model, error) {
	var err error

	switch m.step {
	case 1:
		m, err = StepOne(m)
	}

	if err == nil {
		m.step++
	}

	return m, err
}

func StepOne(m model) (model, error) {
	var err error

	if m.selectedIndex == 0 {
		m.filePath = utils.GetCurrentDirectory()
	} else if m.selectedIndex == 1 {
		m.filePath = m.textInput.Value()
	}

	if m.filePath == "" {
		log.Println("Failed to save path to file. Please try again.")
		err = os.ErrInvalid
	}

	return m, err
}
