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

// Setup calculates the stages with assigned build configurations
func (c *MonoBuild) Setup(limitTo string) error {
	log := c.log.WithField("method", "Setup")
	log.Debugf("setup for %d configurations", len(c.configurations))

	err := c.createStages(c.configurations)
	if err != nil {
		return err
	}

	if len(limitTo) != 0 {
		log.Infof("limiting to configuration [%s]", limitTo)
		for _, stage := range c.stages {
			for _, config := range stage.Configurations {
				if config.Label != limitTo {
					config.skip = true
				}
			}
		}
	}

	// non recursive way of enabling dependencies
	run := true
	for run {
		change := false
		for _, v := range c.configurations {
			if v.skip != true && len(v.Dependencies) > 0 {
				for _, dependency := range v.Dependencies {
					config := c.configurations[dependency]
					if config.skip {
						config.skip = false
						change = true
					}
				}
			}
		}
		if !change {
			run = false
		}
	}

	c.ready = true

	return nil
}
