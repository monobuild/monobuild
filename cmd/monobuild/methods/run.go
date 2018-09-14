package methods

import (
	"github.com/monobuild/monobuild"
	"os"
)

// Run is the wrapper to call the monobuild library
func Run() (returnError error) {
	PrintHeader()

	var err error
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg := monobuild.NewMonoBuild()

	exit := make(chan bool)

	go func() {
		if err = cfg.Run(dir); err != nil {
			returnError = err
		}
		exit <- true
	}()

	<-exit

	return
}
