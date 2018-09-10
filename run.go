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

	for _, stage := range stages {
		log.Debugf("%s", stage)
		for _, cfg := range stage.Configurations {
			log.Debugf("  %s", cfg)
		}
	}

	for _, stage := range stages {
		if err := stage.Execute(); err != nil {
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
