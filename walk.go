package monobuild

import (
	"bytes"
	"os"
	p "path"
	"path/filepath"
	"text/template"

	"github.com/pkg/errors"
)

// Walk iterates through sub directories looking for marker
func (c *Configuration) Walk(baseDir string) error {
	if nil == c.sendChannel {
		return errors.New("no channel provided")
	}
	defer close(c.sendChannel)
	if s, err := os.Stat(baseDir); os.IsNotExist(err) || !s.IsDir() {
		return errors.New("base directory not found")
	}

	tmpl, err := template.New("").Parse(c.Run)
	if err != nil {
		return err
	}

	err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if _, err := os.Stat(p.Join(path, MonoBuildMarker)); err == nil {
				if nil != c.sendChannel {
					var tpl bytes.Buffer
					err := tmpl.Execute(&tpl, path)
					if err != nil {
						return err
					}
					c.sendChannel <- Execute{
						Directory: path,
						Command:   tpl.String(),
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
