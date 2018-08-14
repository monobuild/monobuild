package monobuild

type (
	// MonoBuild contains all data required to run the program
	MonoBuild struct {
		sendChannel chan Execute // sendChannel is used to tell executor to execute other process
	}

	// Execute is sent to executor with relevant information
	Execute struct {
		Directory   string            // Directory is the directory where the marker is found
		Commands    []string          // Command is the command to run
		Environment map[string]string // A list of environment variables to add to the env of the forked process
	}

	// buildConfiguration contains information about how to build the software
	buildConfiguration struct {
		Commands    []string          // Command is the command to run
		Environment map[string]string // A list of environment variables to add to the env of the forked process
	}
)

// NewMonoBuild creates an empty mono build runner
func NewMonoBuild() *MonoBuild {
	return &MonoBuild{}
}
