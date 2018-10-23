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
	"errors"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)

// Run executes the build configurations
func (c *MonoBuild) Run() error {
	log := c.log.WithField("method", "Run")
	log.Debugf("running for %d configurations", len(c.configurations))

	if !c.ready {
		return errors.New("setup not done")
	}

	c.printStageInformation(log)
	c.executeConfigurations(log)

	return nil
}

// executeConfigurations starts building all configurations
func (c *MonoBuild) executeConfigurations(log *logrus.Entry) {
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
}

// printStageInformation prints out some debug information about stages
func (c *MonoBuild) printStageInformation(log *logrus.Entry) {
	for _, stage := range c.stages {
		log.Debugf("%s", stage)
		for _, cfg := range stage.Configurations {
			log.Debugf("  %s", cfg)
		}
	}
}
