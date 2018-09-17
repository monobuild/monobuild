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

import (
	"fmt"
	"github.com/pkg/errors"
)

// createStages builds stages from build configurations
func (c *MonoBuild) createStages(configurations []*BuildConfiguration) error {
	if len(configurations) == 0 {
		return nil
	}
	log := c.log.WithField("method", "createStages")
	var (
		stage *Stage
		err   error
	)
	for len(configurations) > 0 {
		stage, configurations, err = c.createStage(len(c.stages), configurations)
		if err != nil {
			for _, cfg := range configurations {
				log.Warnf("%s could not be added to stage", cfg)
			}
			return err
		}
		stage.Log = c.log
		c.stages = append(c.stages, stage)
	}
	return nil
}

// createStage builds a new stage from build configurations
func (c *MonoBuild) createStage(stageNumber int, configurations []*BuildConfiguration) (*Stage, []*BuildConfiguration, error) {
	log := c.log.WithField("method", "createStages")

	log.Infof("creating `Stage %d`", stageNumber)

	stage := &Stage{
		Label:          fmt.Sprintf("Stage %d", stageNumber),
		Configurations: make([]*BuildConfiguration, 0),
	}

	newConfigurations := make([]*BuildConfiguration, 0)
	before := len(configurations)
	for _, val := range configurations {
		if len(val.Dependencies) == 0 {
			stage.Configurations = append(stage.Configurations, val)
			continue
		}
		add := true
		for _, dep := range val.Dependencies {
			add = add && c.dependencyProcessed(dep)
		}
		if add {
			stage.Configurations = append(stage.Configurations, val)
			continue
		}
		newConfigurations = append(newConfigurations, val)
	}
	after := len(newConfigurations)
	if before == after {
		return nil, newConfigurations, errors.New("dependencies are not valid")
	}
	return stage, newConfigurations, nil
}

// dependencyProcessed checks if a dependency is already build
func (c *MonoBuild) dependencyProcessed(label string) bool {
	for _, stage := range c.stages {
		for _, cfg := range stage.Configurations {
			if cfg.Label == label {
				return true
			}
		}
	}
	return false
}
