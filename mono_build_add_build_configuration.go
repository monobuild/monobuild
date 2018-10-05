package monobuild

import (
	"fmt"
	"github.com/pkg/errors"
)

// AddBuildConfiguration allows adding a build configuration from code
func (c *MonoBuild) AddBuildConfiguration(configuration *BuildConfiguration) error {
	log := c.log.WithField("method", "AddBuildConfiguration")
	log.Debug("adding configuration to list")

	if !configuration.configurationIsValid() {
		return errors.New("configuration is not valid")
	}
	for _, val := range c.configurations {
		if val.Label == configuration.Label {
			return errors.New(fmt.Sprintf("configuration [%s] already exists", val.Label))
		}
	}

	c.configurations = append(c.configurations, configuration)
	return nil
}
