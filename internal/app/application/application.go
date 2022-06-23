package application

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rogpeppe/go-internal/modfile"
	"github.com/saraka/go-tpl/internal/pkg/templates"
)

func Exec(projectPath string) (err error) {
	var (
		input struct {
			ModuleName      string
			ApplicationName string `survey:"application name"`
			ApplicationType string `survey:"application type"`
		}
		modPath = filepath.Join(projectPath, "go.mod")
	)
	mod, err := ioutil.ReadFile(modPath)
	if err != nil {
		return
	}
	input.ModuleName = modfile.ModulePath(mod)
	if err = survey.Ask([]*survey.Question{
		{Name: "application name", Validate: survey.Required, Prompt: &survey.Input{Message: "please input appliction name:"}},
		{Name: "application type", Validate: survey.Required, Prompt: &survey.Select{
			Message: "choose application type:",
			Options: []string{"server", "client"},
		}},
	}, &input); err != nil {
		return
	}
	if err = templates.Generate(input.ApplicationType, projectPath, input); err != nil {
		return
	}
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectPath
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
