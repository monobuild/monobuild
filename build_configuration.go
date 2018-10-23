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

import "fmt"

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
