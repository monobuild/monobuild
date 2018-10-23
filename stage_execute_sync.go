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

import "github.com/hashicorp/go-multierror"

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
