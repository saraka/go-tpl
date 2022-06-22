package initialization

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
)

//go:embed template
var templateFS embed.FS

const (
	rootPath = "template"
	tplExt   = ".tpl"
)

func Init(projectPath string) (err error) {
	var (
		tpl   = template.New(rootPath)
		input struct {
			ProjectName string `survey:"project name"`
			ModuleName  string `survey:"module name"`
		}
	)
	if dir, err := ioutil.ReadDir(projectPath); (err != nil && os.IsExist(err)) || len(dir) > 0 {
		return fmt.Errorf("destination path '%s' already exists and is not an empty directory", projectPath)
	}
	if err = survey.Ask([]*survey.Question{
		{Name: "project name", Validate: survey.Required, Prompt: &survey.Input{Message: "please input project name:"}},
		{Name: "module name", Validate: survey.Required, Prompt: &survey.Input{Message: "please input go module name:"}},
	}, &input); err != nil {
		return
	}
	if err = os.MkdirAll(projectPath, 0755); err != nil {
		return
	}

	cmd := exec.Command("go", "mod", "init", input.ProjectName)
	cmd.Dir = projectPath
	if err = cmd.Run(); err != nil {
		return
	}
	if err = fs.WalkDir(templateFS, rootPath, func(p string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if p == info.Name() {
			return nil
		}
		if tpl, err = tpl.Parse(path.Join(projectPath, strings.TrimPrefix(p, rootPath))); err != nil {
			return err
		}
		var buf bytes.Buffer
		if err = tpl.Execute(&buf, input); err != nil {
			return err
		}
		dst := buf.String()
		if info.IsDir() {
			return os.MkdirAll(dst, 0755)
		}
		if filepath.Ext(dst) == tplExt {
			dst = strings.TrimSuffix(dst, tplExt)
		}
		dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer dstFile.Close()
		if tpl, err = template.ParseFS(templateFS, p); err != nil {
			return err
		}
		if err = tpl.Execute(dstFile, input); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return
	}
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = projectPath
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
