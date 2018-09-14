// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
