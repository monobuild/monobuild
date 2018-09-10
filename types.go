package monobuild

import (
	"github.com/sirupsen/logrus"
)

type (
	// MonoBuild contains all data required to run the program
	MonoBuild struct {
		log *logrus.Entry // base logging facility
	}

	// Execute is sent to executor with relevant information
	Execute struct {
		Directory   string            // Directory is the directory where the marker is found
		Commands    []string          // Command is the command to run
		Environment map[string]string // A list of environment variables to add to the env of the forked process
	}
)

// NewMonoBuild creates an empty mono build runner
func NewMonoBuild() *MonoBuild {
	return &MonoBuild{
		log: logrus.WithField("package", "monobuild"),
	}
}
