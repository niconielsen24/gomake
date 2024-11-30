package ui

import (
	"fmt"
	"gomake/builder"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type preview_model struct {
	title string
	build *builder.BuildArguments
}

func InitialPreviewModel(b *builder.BuildArguments) preview_model {
	return preview_model{
		title: STYLED_TITLE,
		build: b,
	}
}

func (m preview_model) Init() tea.Cmd {
	return nil
}

func (m preview_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m preview_model) View() string {
  s := m.title
  lang := m.build.GetLang()
  techs := m.build.GetTechs()
  
  preview_subtitle_style := lipgloss.NewStyle().
    Bold(true).Padding(0,1,0,1).
    Background(lipgloss.Color("#FFFFFF")).
    Foreground(lipgloss.Color("#000000"))
  var subtitle = preview_subtitle_style.Render("Your current build :")
  s+= fmt.Sprintf("\n%s \n Language : %s\n",subtitle,lang)
  for i:= range len(techs) {
    s += fmt.Sprintf(" | %v : %s\n",i,techs[i])
  } 

	return s
}
