package monobuild

import (
	"github.com/sirupsen/logrus"
)

type (
	// MonoBuild contains all data required to run the program
	MonoBuild struct {
		log *logrus.Entry // base logging facility
	}
)

// NewMonoBuild creates an empty mono build runner
func NewMonoBuild() *MonoBuild {
	return &MonoBuild{
		log: logrus.WithField("package", "monobuild"),
	}
}
