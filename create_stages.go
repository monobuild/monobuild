package monobuild

import (
	"fmt"
	"github.com/pkg/errors"
)

// createStages builds stages from build configurations
func (c *MonoBuild) createStages(configurations []*buildConfiguration) error {
	if len(configurations) == 0 {
		return nil
	}
	log := c.log.WithField("method", "createStages")
	var (
		stage *Stage
		err   error
	)
	for len(configurations) > 0 {
		stage, configurations, err = c.createStage(len(stages), configurations)
		if err != nil {
			for _, cfg := range configurations {
				log.Warnf("%s could not be added to stage", cfg)
			}
			return err
		}
		stage.Log = c.log
		stages = append(stages, stage)
	}
	return nil
}

// createStage builds a new stage from build configurations
func (c *MonoBuild) createStage(stageNumber int, configurations []*buildConfiguration) (*Stage, []*buildConfiguration, error) {
	log := c.log.WithField("method", "createStages")

	log.Infof("creating `Stage %d`", stageNumber)

	stage := &Stage{
		Label:          fmt.Sprintf("Stage %d", stageNumber),
		Configurations: make([]*buildConfiguration, 0),
	}

	newConfigurations := make([]*buildConfiguration, 0)
	before := len(configurations)
	for _, val := range configurations {
		if len(val.Dependencies) == 0 {
			stage.Configurations = append(stage.Configurations, val)
			continue
		}
		add := true
		for _, dep := range val.Dependencies {
			add = add && dependencyProcessed(dep)
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
func dependencyProcessed(label string) bool {
	for _, stage := range stages {
		for _, cfg := range stage.Configurations {
			if cfg.Label == label {
				return true
			}
		}
	}
	return false
}
