package models

import (
	"gomake/builder"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

var builders = map[string]builder.GoMakeBuilder{
	"GO": builder.GoBuilder{},
}

type MsgBuilderStatus int

const (
	BuildOk MsgBuilderStatus = iota
	BuildFail
)

type build_view_model struct {
	config  *builder.BuilderConfig
	builder builder.GoMakeBuilder
	view    viewport.Model
}

func InitialBuildViewModel(bc *builder.BuilderConfig) build_view_model {
	vp := viewport.New(10, 4)
	return build_view_model{
		view:   vp,
		config: bc,
	}
}

func (m build_view_model) Init() tea.Cmd {
	return m.createBuilder
}
func (m build_view_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case MsgBuilderStatus:
		if msg == BuildOk {
      m.builder.InitProj(m.config)
      m.builder.GetDependencies(m.config)
      m.builder.BuildDirs(m.config)
			return m, nil
		}
	}

	return m, nil
}
func (m build_view_model) View() string {
	return ""
}

func (m *build_view_model) createBuilder() tea.Msg {
	if builder, exists := builders[m.config.Lang]; exists {
		m.builder = builder
	} else {
		return BuildFail
	}
	return BuildOk
}
