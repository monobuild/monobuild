package monobuild

import "testing"

func TestCreateStages(t *testing.T) {
	c := NewMonoBuild(".")

	c.AddBuildConfiguration(&BuildConfiguration{
		Label:        "first stage",
		directory:    ".",
		Dependencies: []string{"other"},
		Commands:     []string{"echo ${MY_SECRET_VAR}"},
		Environment: map[string]string{
			"MY_SECRET_VAR": "hello",
		},
	})
	c.AddBuildConfiguration(&BuildConfiguration{
		Label:     "other",
		directory: ".",
		Commands:  []string{"echo other"},
	})
	err := c.Run()
	if err != nil {
		t.FailNow()
	}
	if 2 != len(c.stages) {
		t.Log("there should be two stages")
		t.Fail()
	}
}
