package main

import (
	"fmt"
	"os"

	"github.com/monobuild/monobuild"
	"github.com/monobuild/monobuild/cmd/monobuild/methods"
	"github.com/sirupsen/logrus"
)

// main is the entry method of the monobuild application
func main() {
	// TODO set WarnLevel as default and provide flag to switch
	// TODO provide a quiet flag to suppress header
	logrus.SetLevel(logrus.DebugLevel)
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// run is the wrapper to call the monobuild library
func run() (returnError error) {
	methods.PrintHeader()

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
