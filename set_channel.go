package monobuild

import (
	"github.com/pkg/errors"
)

// SetChannel must be called to allow the logic to send commands to be executed to the executor
func (c *MonoBuild) SetChannel(toExecutor chan Execute) error {
	if nil == toExecutor {
		return errors.New("no channel provided")
	}
	c.sendChannel = toExecutor
	return nil
}
