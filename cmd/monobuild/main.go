package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"livingit.de/code/monobuild"
	"livingit.de/code/monobuild/cmd/monobuild/methods"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() (returnError error) {
	methods.PrintHeader()

	var err error
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg := monobuild.NewMonoBuild()

	rec := make(chan monobuild.Execute)
	cfg.SetChannel(rec)

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
