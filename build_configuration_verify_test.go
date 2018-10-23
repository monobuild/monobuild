package monobuild

import "testing"

func TestInvalidConfiguration(t *testing.T) {
	bc := BuildConfiguration{}
	result := bc.configurationIsValid()
	if result {
		t.Log("passed configuration was invalid but not rejected")
		t.FailNow()
	}
}
