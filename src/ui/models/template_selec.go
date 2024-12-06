package models

import (
	"fmt"
	"gomake/builder/templates"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor       cursor.Model
	current_line int
	options      []string
}

func InitialTemplSelec() model {
	c := cursor.New()
	c.SetChar("->")
	return model{
		cursor:       c,
		current_line: 0,
		options: []string{
			"GO TUI app : bubbletea, bubbles, lipgloss",
			"GO web app : echo, templ, htmx",
			"GO web app : gin",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if m.current_line < len(m.options)-1 {
				m.current_line++
			}
		case "k":
      if m.current_line > 0 {
        m.current_line--
      }
    case " ":
      bc := templates.TuiConfig
      return InitialBuildViewModel(&bc),nil
		}
	}
  nc,cmd := m.cursor.Update(msg)
  m.cursor = nc
	return m, cmd
}

func (m model) View() string {
	s := ""
	for i, option := range m.options {
		if i == m.current_line {
			s += fmt.Sprintf(" %s %s\n", m.cursor.View(), option)
		} else {
			s += fmt.Sprintf("   %s\n", option)
		}
	}
	return s
}
