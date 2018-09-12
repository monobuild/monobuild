package monobuild

import (
	"fmt"
	"os"
)

// environment builds up a new environment variable list for a process to be executed
func (configuration *BuildConfiguration) environment() []string {
	env := os.Environ()
	for k, v := range configuration.Environment {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}
