package ui

import (
	"fmt"
	"gomake/ui/models"

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
	title     string
	status    string
	templates tea.Model
}

func InitialHomeModel() home_model {
	return home_model{
		title:     STYLED_TITLE,
		templates: models.InitialTemplSelec(),
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
		case " ":
		}
	}
  nt, cmd := m.templates.Update(msg)
  m.templates = nt
	return m, cmd
}

func (m home_model) View() string {
	s := m.title + "\n"
	s += fmt.Sprintf("%s", m.templates.View())

	return s
}
