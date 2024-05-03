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
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

// BuildConfiguration contains information about how to build the software
type BuildConfiguration struct {
	Commands     []string          `yaml:"commands"`                  // Command is the command to run
	Environment  map[string]string `yaml:"environment"`               // A list of environment variables to add to the env of the forked process
	Label        string            `yaml:"label" validate:"required"` // Label is the name of the build configuration
	Dependencies []string          `yaml:"dependencies"`              // A list of dependencies to other build configurations
	Parallel     bool              `yaml:"parallel"`                  // Parallel determines whether build configuration is allowed to run in parallel

	directory string `yaml:"-"` // directory is used to store the directory of the build configuration
	skip      bool   `yaml:"-"` // skip is a flag to skip configuration based on the limit flag
}

// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (configuration *BuildConfiguration) String() string {
	return fmt.Sprintf("build configuration `%s`", configuration.Label)
}

func (configuration *BuildConfiguration) configurationIsValid() bool {
	validate := validator.New()

	err := validate.Struct(configuration)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logrus.Errorf("%s is %s", err.StructField(), err.ActualTag())
		}
		return false
	}
	return true
}

// SetDirectory stores the directory of the build configuration
func (configuration *BuildConfiguration) SetDirectory(path string) {
	configuration.directory = path
}

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
