package monobuild

import (
	"github.com/pkg/errors"
	"os"
	p "path"
	"path/filepath"
)

// walk iterates through sub directories looking for marker
func (c *MonoBuild) walk(baseDir string) ([]*buildConfiguration, error) {
	log := c.log.WithField("method", "walk")
	configs := make([]*buildConfiguration, 0)
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			marker := p.Join(path, MonoBuildMarker)
			if _, err := os.Stat(marker); err == nil {
				log.Infof("build configuration found at %s", marker)
				bc, err := c.loadBuildConfiguration(marker)
				if err != nil {
					return err
				}
				if !bc.configurationIsValid() {
					return errors.New("build configuration is not valid")
				}
				bc.Directory = path
				configs = append(configs, bc)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return configs, nil
}
