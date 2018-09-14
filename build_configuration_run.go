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
	"github.com/hashicorp/go-multierror"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
	"os"
	"strings"
)

// run executes a single configuration of a stage
func (configuration *BuildConfiguration) run(stage *Stage) *multierror.Error {
	var (
		result *multierror.Error
	)
	log := stage.Log.WithField("method", "run")

	for _, cmd := range configuration.Commands {
		log.Debugf("about to run %s", cmd)
		p, err := syntax.NewParser().Parse(strings.NewReader(cmd), "")
		if err != nil {
			result = multierror.Append(result, err)
		}

		env := configuration.environment()
		r := interp.Runner{
			Dir: configuration.directory,
			Env: env,

			Exec: interp.DefaultExec,
			Open: interp.OpenDevImpls(interp.DefaultOpen),

			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
		if err = r.Reset(); err != nil {
			result = multierror.Append(result, err)
		}
		err = r.Run(p)
		if err != nil {
			result = multierror.Append(result, err)
		}
		if nil != result {
			break
		}
	}
	return result
}
