package main

import (
	"fmt"
	"gomake/ui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  m := ui.InitialHomeModel()
  p := tea.NewProgram(m,tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Fuck! something broke : %s", err.Error())
    os.Exit(1)
  } 
}
