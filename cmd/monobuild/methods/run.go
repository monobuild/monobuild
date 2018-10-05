package methods

import (
	"github.com/monobuild/monobuild"
	"github.com/spf13/viper"
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

	cfg.DisableParallelism = viper.GetBool("no-parallelism")
	cfg.MarkerName = viper.GetString("marker")

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
