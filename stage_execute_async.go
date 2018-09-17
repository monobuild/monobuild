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
	"sync"
)

// executeAsync runs all configurations that are allowed to run in parallel
func (stage *Stage) executeAsync() {
	log := stage.Log.WithField("method", "executeAsync")

	var wg sync.WaitGroup
	for _, cfg := range stage.Configurations {
		if cfg.Parallel {
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
		log.Errorf("error asynchronously running configuration: %s", err.Error())
	}
	log.Debugf("parallel execution done: %s", configuration)
}
