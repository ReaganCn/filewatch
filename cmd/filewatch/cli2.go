package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Uses bubbletea to create a CLI
/* Create a model to track the state of our application */
type model struct {
	textInput     textinput.Model
	choices       []string // items on the to-do list
	cursor        int      // which to-do list item our cursor is pointing at
	selected      string   // which to-do items are selected
	selectedIndex int      // the index of the selected item
}

func bubble() {
	p := tea.NewProgram(initialModel())

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.selected != "" {
		fmt.Printf("\n---\nYou chose: %s!\n", m.selected)

		result := m.fileMonitorChoice(m.selectedIndex)

		fmt.Println("Result: ", result)

		// Save the result to the database
	}
}

func initialModel() model {
	ti := textinput.New()
	ti.Focus()

	return model{

		// Initialize the text input
		textInput: ti,
		// Our to-do list is a grocery list
		choices: []string{"Current directory", "Enter file path"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: "",
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			m.selected = m.choices[m.cursor]
			m.selectedIndex = m.cursor
			return m, tea.Quit
		}

		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {

	s := strings.Builder{}

	// The header
	s.WriteString(subtitleStyle.Render("We will guide you through the process of setting up filewatch. 🏎️"))
	s.WriteString("\nWhat do you want to monitor?\n\n")

	// Iterate over our choices
	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("[x] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}

	if m.cursor == 1 {
		s.WriteString("\n")
		s.WriteString(m.textInput.View())
	}

	s.WriteString(quitStyle.Render("\n(press q to quit)\n"))

	return s.String()
}

func (m model) fileMonitorChoice(index int) string {

	switch index {
	case 0:
		return getCurrentDirectory()
	case 1:
		return ""
	default:
		return ""
	}
}

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dir
}
