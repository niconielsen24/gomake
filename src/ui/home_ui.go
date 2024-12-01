package ui

import (
	"fmt"
	"gomake/builder"
	"math"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
)

const TITLE = `
    _______   ______   .___  ___.      ___       __  ___  _______    
   /  _____| /  __  \  |   \/   |     /   \     |  |/  / |   ____|  
  |  |  __  |  |  |  | |  \  /  |    /  ^  \    |  '  /  |  |__     
  |  | |_ | |  |  |  | |  |\/|  |   /  /_\  \   |    <   |   __|    
  |  |__| | |  '--'  | |  |  |  |  /  _____  \  |  .  \  |  |____   
   \______|  \______/  |__|  |__| /__/     \__\ |__|\__\ |_______|  
  `

var STYLED_TITLE = Title_style.Render(TITLE)

type selections map[int]rune

type home_model struct {
	title                 string
	options               Options
	selections            selections
	cursor                cursor.Model
	current_line          int
	current_selected_line int
	current_build         *builder.BuildArguments
	current_selection     string
	is_initial            bool
}

func InitialHomeModel() home_model {
	c := cursor.New()
	c.Blink = true
	c.SetChar("->")
	s := make(selections)
	do := GetDefaultOptions()
	for i := 0; i < len(do); i++ {
		s[i] = ' '
	}
	return home_model{
		title:             STYLED_TITLE,
		options:           do,
		selections:        s,
		cursor:            c,
		current_line:      0,
		current_build:     nil,
		current_selection: "",
		is_initial:        true,
	}
}

func (m home_model) Init() tea.Cmd {
	return nil
}

func (m home_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "down", "j":
			if m.current_line < len(m.options)-1 {
				m.current_line++
			}
		case "up", "k":
			if m.current_line > 0 {
				m.current_line--
			}
		case "r", "backspace":
			return InitialHomeModel(), nil
		case "enter", " ":
			if m.is_initial {
				for i := range len(m.selections) {
					if m.selections[i] == 'X' {
						m.selections[i] = ' '
					}
				}
				m.selections[m.current_line] = 'X'
				m.current_selected_line = m.current_line
				m.current_selection = m.options[m.current_line]
			} else {
				if m.selections[m.current_line] == 'X' {
					m.selections[m.current_line] = ' '
					return m, nil
				}
				m.selections[m.current_line] = 'X'
			}
		case "a":
			if m.is_initial {
				if m.current_selection != "" {
					ba, err := builder.ParseArguments(m.current_selection)
					if err != nil {
						return m, tea.Quit
					}
					m.is_initial = false
					m.current_build = ba
					m.current_line = 0
					m.selections[m.current_selected_line] = ' '
					m.current_selected_line = math.MaxInt
					m.options = GetLangOptions(ba.GetLang())
					return m, nil
				}
			} else {
				var selected_techs []string
				for i := range len(m.options) {
					if m.selections[i] == 'X' {
						selected_techs = append(selected_techs, m.options[i])
					}
				}
				if len(selected_techs) != 0 {
					m.current_build.SetTechs(selected_techs)
					return InitialPreviewModel(m.current_build), nil
				}
			}
		}
	}
	return m, nil
}

func (m home_model) View() string {
	s := m.title + "\n"

	for i := range len(m.options) {
		if m.current_line == i {
			s += fmt.Sprintf(" %v [%c] %v\n", m.cursor.View(), m.selections[i], m.options[i])
		} else {
			s += fmt.Sprintf("   [%c] %v\n", m.selections[i], m.options[i])
		}
	}

	s += `

  Press q to quit
  Press [enter, space] to select current line
  Press a to accept current selection
  `
	if !m.is_initial {
		s += "Press [r, bckspace] to go back to the last panel\n"
	}

	return s
}
