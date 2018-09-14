package monobuild

// Execute runs the content of a stage
func (stage *Stage) Execute() error {
	result := stage.executeSync()
	if nil == result {
		stage.executeAsync()
	}

	return result.ErrorOrNil()
}
