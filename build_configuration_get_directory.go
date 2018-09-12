package monobuild

// GetConfiguration returns the path where the build configuration is located
func (configuration *BuildConfiguration) GetDirectory() string {
	return configuration.directory
}
