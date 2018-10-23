package monobuild

import "testing"

func TestConfigurationToSkip(t *testing.T) {
	c := NewMonoBuild(".")

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:     "first stage",
		directory: ".",
		Commands:  []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
		skip: true,
	})

	/* this is a call to a private method
	 * therefore this does not classify as a black box test for
	 * the api */
	if err := c.createStages(c.configurations); err != nil {
		t.Log("There should be stages")
		t.FailNow()
	}

	if 1 != len(c.stages) {
		t.Log("there should be one stage")
		t.FailNow()
	}

	if err := c.stages[0].Configurations[0].run(c.stages[0]); err != nil {
		t.Log("there should be no error because the build configuration should be skipped")
		t.Logf("error returned: %s", err)
		t.FailNow()
	}
}

func TestSkipNoDependencies(t *testing.T) {
	c := NewMonoBuild(".")

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:     "first stage",
		directory: ".",
		Commands:  []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:     "main",
		directory: ".",
		Commands:  []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	if err := c.Setup("main"); err != nil {
		t.Logf("setup returned error: %s", err)
		t.FailNow()
	}

	numberOfConfigurations := 0
	for _, stage := range c.stages {
		for _, config := range stage.Configurations {
			if !config.skip {
				numberOfConfigurations++
			}
		}
	}

	if numberOfConfigurations != 1 {
		t.Log("only one configuration should be active")
		t.FailNow()
	}
}

func TestSkipWithDependency(t *testing.T) {
	c := NewMonoBuild(".")

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:     "first stage",
		directory: ".",
		Commands:  []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:        "main",
		directory:    ".",
		Dependencies: []string{"sub"},
		Commands:     []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:     "sub",
		directory: ".",
		Commands:  []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	if err := c.Setup("main"); err != nil {
		t.Logf("setup returned error: %s", err)
		t.FailNow()
	}

	numberOfConfigurations := 0
	for _, stage := range c.stages {
		for _, config := range stage.Configurations {
			if !config.skip {
				numberOfConfigurations++
			}
		}
	}

	if numberOfConfigurations != 2 {
		t.Log("dependencies should be executed")
		t.FailNow()
	}
}

func TestSkipWithIndirectDependency(t *testing.T) {
	c := NewMonoBuild(".")

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:     "first stage",
		directory: ".",
		Commands:  []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:        "main",
		directory:    ".",
		Dependencies: []string{"sub"},
		Commands:     []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:        "sub",
		directory:    ".",
		Commands:     []string{"echo ${MY_SECRET_VAR}"},
		Dependencies: []string{"sub-sub"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:     "sub-sub",
		directory: ".",
		Commands:  []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})

	if err := c.Setup("main"); err != nil {
		t.Logf("setup returned error: %s", err)
		t.FailNow()
	}

	numberOfConfigurations := 0
	for _, stage := range c.stages {
		for _, config := range stage.Configurations {
			if !config.skip {
				numberOfConfigurations++
			}
		}
	}

	if numberOfConfigurations != 3 {
		t.Log("indirect dependencies should be executed")
		t.FailNow()
	}
}
