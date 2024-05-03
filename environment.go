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
	"os"
	"strings"

	"mvdan.cc/sh/expand"
)

type MonobuildEnviron struct {
	environmentVariables map[string]expand.Variable
}

func newMonobuildEnviron() MonobuildEnviron {
	return MonobuildEnviron{environmentVariables: map[string]expand.Variable{}}
}

func (m MonobuildEnviron) Get(name string) expand.Variable {
	if v, ok := m.environmentVariables[name]; ok {
		return v
	}
	return expand.Variable{}
}

func (m MonobuildEnviron) Each(f func(name string, vr expand.Variable) bool) {
	for s := range m.environmentVariables {
		if !f(s, m.environmentVariables[s]) {
			break
		}
	}
}

func (m MonobuildEnviron) Set(name string, vr expand.Variable) {
	m.environmentVariables[name] = vr
}

// environment builds up a new environment variable list for a process to be executed
func (configuration *BuildConfiguration) environment() (expand.Environ, error) {
	var m = newMonobuildEnviron()
	env := os.Environ()
	for i := range env {
		kv := strings.SplitN(env[i], "=", 2)
		m.Set(kv[0], expand.Variable{
			Local:    false,
			Exported: true,
			ReadOnly: true,
			NameRef:  true,
			Value:    kv[1],
		})
	}
	for k, v := range configuration.Environment {
		m.Set(k, expand.Variable{
			Local:    false,
			Exported: true,
			ReadOnly: true,
			NameRef:  true,
			Value:    v,
		})
	}
	return m, nil
}
