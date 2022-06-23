package initialization

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/saraka/go-tpl/internal/app/application"
	"github.com/saraka/go-tpl/internal/pkg/templates"
)

func Exec(projectPath string) (err error) {
	var (
		input struct {
			ModuleName string `survey:"module name"`
		}
	)
	if dir, err := ioutil.ReadDir(projectPath); (err != nil && os.IsExist(err)) || len(dir) > 0 {
		return fmt.Errorf("destination path '%s' already exists and is not an empty directory", projectPath)
	}
	if err = survey.Ask([]*survey.Question{
		{Name: "module name", Validate: survey.Required, Prompt: &survey.Input{Message: "please input go module name:"}},
	}, &input); err != nil {
		return
	}
	if err = os.MkdirAll(projectPath, 0755); err != nil {
		return
	}

	cmd := exec.Command("go", "mod", "init", input.ModuleName)
	cmd.Dir = projectPath
	if err = cmd.Run(); err != nil {
		return
	}

	if err = templates.Generate("init", projectPath, input); err != nil {
		return
	}

	if err = application.Exec(projectPath); err != nil {
		return
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = projectPath
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
