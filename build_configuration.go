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
}

// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (configuration *BuildConfiguration) String() string {
	return fmt.Sprintf("build configuration `%s`", configuration.Label)
}
