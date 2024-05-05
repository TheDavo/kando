package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))

	todoStatusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#990000")).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderBottom(true)
	inProgStatuStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFF00")).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderBottom(true)
	doneStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#009900")).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderBottom(true)
)
