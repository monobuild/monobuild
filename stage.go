package monobuild

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// Stage is a build stage containing only independent build configurations
type Stage struct {
	Label          string
	Configurations []*buildConfiguration
	Log            *logrus.Entry
}

// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (stage *Stage) String() string {
	return fmt.Sprintf("%s - %d build configurations", stage.Label, len(stage.Configurations))
}
