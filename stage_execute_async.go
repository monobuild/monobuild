package monobuild

import (
	"sync"
)

// executeAsync runs all configurations that are allowed to run in parallel
func (stage *Stage) executeAsync() {
	log := stage.Log.WithField("method", "executeAsync")

	var wg sync.WaitGroup
	for _, cfg := range stage.Configurations {
		if !cfg.Parallel {
			wg.Add(1)
			log.Infof("parallel working on %s", cfg)
			go stage.executeWithWg(cfg, &wg)
		}
	}
	wg.Wait()
}

// executeWithWg is a wrapper to run build asynchronously
func (stage *Stage) executeWithWg(configuration *BuildConfiguration, wg *sync.WaitGroup) {
	log := stage.Log.WithField("method", "executeAsync")
	defer wg.Done()

	err := configuration.run(stage)
	if nil != err {
		log.Errorf("error asynchronously running configuration: %s", err.Error)
	}
	log.Debugf("parallel execution done: %s", configuration)
}
