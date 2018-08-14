package main

import (
	"fmt"
	"os"

	"livingit.de/code/monobuild"
	"livingit.de/code/monobuild/cmd/monobuild/methods"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
	"strings"
)

func main() {
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
		if err = cfg.Walk(dir); err != nil {
			returnError = err
		}
		exit <- true
	}()

	for i := range rec {
		fmt.Printf("%#v\n", i)
		for _, cmd := range i.Commands {
			fmt.Println(cmd)
			p, err := syntax.NewParser().Parse(strings.NewReader(cmd), "")
			if err != nil {
				return err
			}

			env := os.Environ()
			for k, v := range i.Environment {
				env = append(env, fmt.Sprintf("%s=%s", k, v))
			}

			r := interp.Runner{
				Dir: i.Directory,
				Env: env,

				Exec: interp.DefaultExec,
				Open: interp.OpenDevImpls(interp.DefaultOpen),

				Stdin:  os.Stdin,
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}
			if err = r.Reset(); err != nil {
				return err
			}
			r.Run(p)
		}
	}

	<-exit

	return
}
