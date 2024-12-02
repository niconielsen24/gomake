package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	title string
	left  viewport.Model
	right viewport.Model
}

func InitialHomeModel() home_model {
  left := viewport.New(5,2)
  left.Style = lipgloss.NewStyle().Border(lipgloss.DoubleBorder(),true)
  right := viewport.New(5,2)
  right.Style = lipgloss.NewStyle().Border(lipgloss.DoubleBorder(),true)

	return home_model{
		title: STYLED_TITLE,
		left:  left,
		right: right,
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
		}
	}
	return m, nil
}

func (m home_model) View() string {
	s := m.title + "\n"
	s += fmt.Sprintf("%s%s", m.left.View(), m.right.View())

	return s
}
