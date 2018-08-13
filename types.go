package monobuild

type (
	// Configuration contains all data required to run the program
	Configuration struct {
		Label       string            `yaml:"label"`    // Label is the name of the file to look for
		Run         string            `yaml:"run"`      // Run is executed in each labelled directory
		Environment map[string]string `yaml:"env-vars"` // Environment contains a list of environment variables to set

		sendChannel chan Execute // sendChannel is used to tell executor to execute other process
	}

	// Execute is sent to executor with relevant information
	Execute struct {
		Directory string
		Command   string
	}
)
