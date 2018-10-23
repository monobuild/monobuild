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
	"mvdan.cc/sh/interp"
	"os"
)

// environment builds up a new environment variable list for a process to be executed
func (configuration *BuildConfiguration) environment() (interp.Environ, error) {
	env := os.Environ()
	for k, v := range configuration.Environment {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return interp.EnvFromList(env)
}
