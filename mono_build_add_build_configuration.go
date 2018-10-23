package monobuild

import (
	"github.com/pkg/errors"
)

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
