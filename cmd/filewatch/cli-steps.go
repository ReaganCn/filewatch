package main

import (
	"log"
	"os"

	"github.com/reagancn/filewatch/pkg/database/models"
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

	p := models.Path{Path: m.filePath}

	if m.filePath == "" {
		log.Println("Failed to save path to file. Please try again.")
		err = os.ErrInvalid
	} else {
		id, err := p.Insert(m.dbConn)

		if err != nil {
			log.Println("Failed to save path to file. Please try again.")
		} else {
			log.Println("Successfully saved path to file. ID: ", id)
		}
	}

	return m, err
}
