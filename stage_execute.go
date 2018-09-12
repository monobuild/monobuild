package monobuild

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
	"os"
	"strings"
)

// Execute runs the content of a stage
func (stage *Stage) Execute() error {
	var (
		result *multierror.Error
	)
	log := stage.Log.WithField("method", "Execute")

	// TODO decide on parallelism
	for _, cfg := range stage.Configurations {
		if !cfg.Parallel {
			log.Infof("working on %s", cfg)
			result = stage.runConfiguration(cfg, log, result)
			if nil != result {
				log.Error("error running configuration, breaking")
				break
			}
		}
	}
	return result.ErrorOrNil()
}

// runConfiguration executes a single configuration of a stage
func (stage *Stage) runConfiguration(cfg *BuildConfiguration, log *logrus.Entry, result *multierror.Error) *multierror.Error {
	for _, cmd := range cfg.Commands {
		log.Debugf("about to run %s", cmd)
		p, err := syntax.NewParser().Parse(strings.NewReader(cmd), "")
		if err != nil {
			result = multierror.Append(result, err)
		}

		env := buildEnvironment(cfg)
		r := interp.Runner{
			Dir: cfg.directory,
			Env: env,

			Exec: interp.DefaultExec,
			Open: interp.OpenDevImpls(interp.DefaultOpen),

			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
		if err = r.Reset(); err != nil {
			result = multierror.Append(result, err)
		}
		err = r.Run(p)
		if err != nil {
			result = multierror.Append(result, err)
		}
		if nil != result {
			break
		}
	}
	return result
}

// buildEnvironment builds up a new environment variable list for a process to be executed
func buildEnvironment(cfg *BuildConfiguration) []string {
	env := os.Environ()
	for k, v := range cfg.Environment {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}
