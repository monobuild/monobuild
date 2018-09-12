package monobuild

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func (c *MonoBuild) loadBuildConfiguration(configurationFile string) (*BuildConfiguration, error) {
	var bc BuildConfiguration

	content, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &bc)
	if err != nil {
		return nil, err
	}

	return &bc, nil
}
