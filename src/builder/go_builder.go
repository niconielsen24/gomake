package builder

import (
	"os/exec"
)

type GoBuilder struct {
	techs []string
}

var techs = map[string]string{
  "bubbletea" : "",
  "bubbles" : "",
  "lipgloss" : "",
  "echo" : "",
}

func (gb *GoBuilder) Build(bc *BuildConfig) error {
  cmd := exec.Command("go", "mod init")
  if err := cmd.Run() ; err != nil {
    return err
  }
  return nil
}

