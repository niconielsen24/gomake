package builder

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type GoBuilder struct {
	techs []string
}

var techs = map[string]string{
	"bubbletea": "github.com/charmbracelet/bubbletea",
	"bubbles":   "github.com/charmbracelet/bubbles",
	"lipgloss":  "github.com/charmbracelet/lipgloss",
	"echo":      "github.com/labstack/echo",
	"gin":       "github.com/gin-gonic/gin",
	"templ":     "github.com/a-h/templ",
}

func (gb GoBuilder) InitProj(bc *BuilderConfig) (string, error) {
	cmd := exec.Command("go", "mod", "init", bc.ProjectName)
	return runCmd(cmd)
}

func (gb GoBuilder) GetDependencies(bc *BuilderConfig) (string, error) {
	var dep_stats string
	for _, dep := range bc.Techs {
		if path, exists := techs[dep]; exists {
			cmd := exec.Command("go", "get", path)
			if stat, err := runCmd(cmd); err != nil {
				return dep_stats + "\nError : " + err.Error(), err
			} else {
				dep_stats += stat
			}
		}
	}
	return dep_stats, nil
}

func (gb GoBuilder) BuildDirs(bc *BuilderConfig) error {
  cwd, _ := os.Getwd()
  bc.DirTree.BuildDirTree(cwd)
	return nil
}

func runCmd(cmd *exec.Cmd) (string, error) {
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		return "Error : " + err.Error(), err
	}

	o, _ := io.ReadAll(stdout)
	e, _ := io.ReadAll(stderr)

	cmd.Wait()
	exit_code := cmd.ProcessState.ExitCode()
	str := fmt.Sprintf("%s\n%s\nExit code %v", o, e, exit_code)
	return str, nil
}
