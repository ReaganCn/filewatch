package main

import "github.com/reagancn/filewatch/pkg/utils"

func StepsHandler(m model) model {
	switch m.step {
	case 1:
		m = StepOne(m)
	}
	m.step++

	return m
}

func StepOne(m model) model {

	if m.selectedIndex == 0 {
		m.filePath = utils.GetCurrentDirectory()
	} else if m.selectedIndex == 1 {
		m.filePath = m.textInput.Value()
	}
	return m
}
