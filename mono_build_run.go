// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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

import "github.com/hashicorp/go-multierror"

func (c *MonoBuild) Run(baseDir string) error {
	log := c.log.WithField("method", "run")
	configurations, err := c.walk(baseDir)
	if err != nil {
		return err
	}

	err = c.createStages(configurations)
	if err != nil {
		return err
	}

	for _, stage := range c.stages {
		log.Debugf("%s", stage)
		for _, cfg := range stage.Configurations {
			log.Debugf("  %s", cfg)
		}
	}

	for _, stage := range c.stages {
		if err := stage.Execute(c.DisableParallelism); err != nil {
			if multiError, ok := err.(*multierror.Error); ok {
				for _, err := range multiError.Errors {
					log.Errorf("could not run stage %s: %s", stage, err)
				}
			}
			break
		}
	}

	return nil
}
