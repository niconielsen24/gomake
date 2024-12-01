package ui

import (
	"fmt"
	"gomake/builder"
	"os"

	tea "github.com/charmbracelet/bubbletea"
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
    case "enter", " " :
      //THIS HAS TO GO
      _,_ = os.Create("file")
    }
	}
	return m, nil
}

func (m preview_model) View() string {
  s := m.title
  lang := m.build.GetLang()
  techs := m.build.GetTechs()
  
  var subtitle = Subtitle_style.Render("Your current build :")
  s+= fmt.Sprintf("\n%s \n Language : %s\n",subtitle,lang)
  for i:= range len(techs) {
    s += fmt.Sprintf(" | %v : %s\n",i,techs[i])
  } 

	return s
}
