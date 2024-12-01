package ui

import(
  "github.com/charmbracelet/lipgloss"
)

var Title_style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4")).
	PaddingBottom(1).
	Border(lipgloss.RoundedBorder(), true)

var Subtitle_style = lipgloss.NewStyle().
    Bold(true).Padding(0,1,0,1).
    Background(lipgloss.Color("#FFFFFF")).
    Foreground(lipgloss.Color("#000000"))
