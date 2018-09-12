package monobuild

import "github.com/hashicorp/go-multierror"

// executeSync runs all configurations that are not allowed to run in parallel
func (stage *Stage) executeSync() *multierror.Error {
	var (
		result *multierror.Error
	)
	log := stage.Log.WithField("method", "executeSync")

	for _, cfg := range stage.Configurations {
		if !cfg.Parallel {
			log.Infof("working on %s", cfg)
			err := cfg.run(stage)
			if nil != err {
				result = multierror.Append(result, err)
				log.Error("error running configuration, breaking")
				break
			}
		}
	}
	return result
}
