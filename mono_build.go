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
	"fmt"
	"os"
	p "path"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type (
	// MonoBuild contains all data required to run the program
	MonoBuild struct {
		log            *logrus.Entry                  // base logging facility
		stages         []*Stage                       // build stages are stored here
		baseDirectory  string                         // baseDirectory determines the root from where to scan
		configurations map[string]*BuildConfiguration // configurations all found configurations
		ready          bool                           // ready denotes whether the setup is done

		DisableParallelism bool   // DisableParallelism can be used to run build independently of markers
		MarkerName         string // MarkerName is the name of the file with the build configuration
	}
)

// NewMonoBuild creates an empty mono build runner
func NewMonoBuild(baseDir string) *MonoBuild {
	return &MonoBuild{
		log:            logrus.WithField("package", "monobuild"),
		baseDirectory:  baseDir,
		configurations: make(map[string]*BuildConfiguration),
	}
}

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

// LoadConfigurations scans the directory tree for build configurations
func (c *MonoBuild) LoadConfigurations() error {
	c.log.WithField("method", "run").Debug("Loading configurations from filesystem")
	return c.walk()
}

func (c *MonoBuild) loadBuildConfiguration(configurationFile string) (*BuildConfiguration, error) {
	var bc BuildConfiguration

	content, err := os.ReadFile(configurationFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &bc)
	if err != nil {
		return nil, err
	}

	return &bc, nil
}

// createStages builds stages from build configurations
func (c *MonoBuild) createStages(configurations map[string]*BuildConfiguration) error {
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
func (c *MonoBuild) createStage(stageNumber int, configurations map[string]*BuildConfiguration) (*Stage, map[string]*BuildConfiguration, error) {
	log := c.log.WithField("method", "createStages")

	log.Infof("creating `Stage %d`", stageNumber)

	stage := &Stage{
		Label:          fmt.Sprintf("Stage %d", stageNumber),
		Configurations: make([]*BuildConfiguration, 0),
	}

	newConfigurations := make(map[string]*BuildConfiguration, 0)
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
		newConfigurations[val.Label] = val
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

// AddBuildConfiguration allows adding a build configuration from code
func (c *MonoBuild) AddBuildConfiguration(configuration *BuildConfiguration) error {
	log := c.log.WithField("method", "AddBuildConfiguration")
	log.Debug("adding configuration to list")

	if !configuration.configurationIsValid() {
		return errors.New("configuration is not valid")
	}

	if _, ok := c.configurations[configuration.Label]; ok {
		return errors.New("non unique Labels found")
	}

	c.configurations[configuration.Label] = configuration
	return nil
}
