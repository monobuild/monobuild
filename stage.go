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
	"sync"

	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)

// Stage is a build stage containing only independent build configurations
type Stage struct {
	Label          string
	Configurations []*BuildConfiguration
	Log            *logrus.Entry
}

// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (stage *Stage) String() string {
	return fmt.Sprintf("%s - %d build configurations", stage.Label, len(stage.Configurations))
}

// Execute runs the content of a stage
func (stage *Stage) Execute(disableParallelism bool) error {
	result := stage.executeSync(disableParallelism)
	if nil == result && !disableParallelism {
		stage.executeAsync()
	}

	return result.ErrorOrNil()
}

// executeSync runs all configurations that are not allowed to run in parallel
func (stage *Stage) executeSync(all bool) *multierror.Error {
	var (
		result *multierror.Error
	)
	log := stage.Log.WithField("method", "executeSync")

	for _, cfg := range stage.Configurations {
		if !cfg.Parallel || all {
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
