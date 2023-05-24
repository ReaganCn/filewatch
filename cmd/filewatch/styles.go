package main

import "github.com/charmbracelet/lipgloss"

var quitStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Width(22)

var subtitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4B0082"))

var titleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#4B0082"))
