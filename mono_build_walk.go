// Copyright © 2018 Sascha Andres <sascha.andres@outlook.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monobuild

import (
	"github.com/pkg/errors"
	"os"
	p "path"
	"path/filepath"
)

// walk iterates through sub directories looking for marker
func (c *MonoBuild) walk() error {
	log := c.log.WithField("method", "walk")
	err := filepath.Walk(c.baseDirectory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			marker := p.Join(path, c.MarkerName)
			if _, err := os.Stat(marker); err == nil {
				log.Infof("build configuration found at %s", marker)
				bc, err := c.loadBuildConfiguration(marker)
				if err != nil {
					return err
				}
				if !bc.configurationIsValid() {
					return errors.New("build configuration is not valid")
				}
				bc.SetDirectory(path)
				if err := c.AddBuildConfiguration(bc); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}
