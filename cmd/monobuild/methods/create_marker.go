package methods

import (
	"os"
	"text/template"
)

func CreateMarkerFile(markerFileName string) error {
	directory, err := os.Getwd()
	if err != nil {
		return err
	}
	directoryInfo, err := os.Stat(directory)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(markerTemplate)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(markerFileName, os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, struct {
		Directory string
		Version   string
	}{
		directoryInfo.Name(),
		versionNumber,
	})
}
