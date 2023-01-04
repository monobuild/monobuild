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
	"context"
	"os"
	"strings"

	"github.com/hashicorp/go-multierror"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

// run executes a single configuration of a stage
func (configuration *BuildConfiguration) run(stage *Stage) *multierror.Error {
	var (
		result *multierror.Error
	)
	log := stage.Log.WithField("method", "run")

	if configuration.skip {
		log.Infof("%s will be skipped", configuration)
		return nil
	}

	for _, cmd := range configuration.Commands {
		log.Debugf("about to run %s", cmd)
		p, err := syntax.NewParser().Parse(strings.NewReader(cmd), "")
		if err != nil {
			result = multierror.Append(result, err)
		}

		env, err := configuration.environment()
		if err != nil {
			result = multierror.Append(result, err)
		}
		r, err := interp.New(interp.Env(env), interp.StdIO(os.Stdin, os.Stdout, os.Stderr), interp.Dir(configuration.directory))
		if err != nil {
			result = multierror.Append(result, err)
			return result
		}
		r.Reset()
		err = r.Run(context.Background(), p)
		if err != nil {
			result = multierror.Append(result, err)
		}
		if nil != result {
			break
		}
	}
	return result
}
