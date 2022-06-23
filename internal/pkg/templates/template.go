package templates

import (
	"bytes"
	"embed"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed init
//go:embed client
//go:embed server
var FS embed.FS

var (
	tpl    = template.New("")
	tplExt = ".tpl"
)

func Generate(tplPath, projectPath string, data interface{}) error {
	return fs.WalkDir(FS, tplPath, func(p string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if p == info.Name() {
			return nil
		}
		if tpl, err = tpl.Parse(path.Join(projectPath, strings.TrimPrefix(p, tplPath))); err != nil {
			return err
		}
		var buf bytes.Buffer
		if err = tpl.Execute(&buf, data); err != nil {
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
		if tpl, err = template.ParseFS(FS, p); err != nil {
			return err
		}
		if err = tpl.Execute(dstFile, data); err != nil {
			return err
		}
		return nil
	})
}
