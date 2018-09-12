package monobuild

import "github.com/hashicorp/go-multierror"

// executeAsync runs all configurations that are allowed to run in parallel
func (stage *Stage) executeAsync() *multierror.Error {
	var (
		result *multierror.Error
	)
	log := stage.Log.WithField("method", "executeAsync")

	for _, cfg := range stage.Configurations {
		if !cfg.Parallel {
			log.Infof("parallel working on %s", cfg)
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
