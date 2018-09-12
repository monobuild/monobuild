package monobuild

// SetDirectory stores the directory of the build configuration
func (configuration *BuildConfiguration) SetDirectory(path string) {
	configuration.directory = path
}
