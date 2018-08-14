package monobuild

import (
	"github.com/pkg/errors"
	"os"
	p "path"
	"path/filepath"
)

// Walk iterates through sub directories looking for marker
func (c *MonoBuild) Walk(baseDir string) error {
	if nil == c.sendChannel {
		return errors.New("no channel provided")
	}
	defer close(c.sendChannel)
	if s, err := os.Stat(baseDir); os.IsNotExist(err) || !s.IsDir() {
		return errors.New("base directory not found")
	}

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if _, err := os.Stat(p.Join(path, MonoBuildMarker)); err == nil {
				bc, err := c.loadBuildConfiguration(p.Join(path, MonoBuildMarker))
				if err != nil {
					return err
				}
				if nil != c.sendChannel {
					c.sendChannel <- Execute{
						Directory:   path,
						Commands:    bc.Commands,
						Environment: bc.Environment,
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
