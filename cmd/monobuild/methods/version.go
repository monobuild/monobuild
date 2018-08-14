package methods

import (
	"fmt"
	"os"
)

var (
	versionNumber = "develop"
	commit        = "dirty"
	date          = "n/a"
)

func PrintHeader() {
	fmt.Fprintln(os.Stdout, "monobuild")
	fmt.Fprintf(os.Stdout, "%s ( %s ) build on %s\n", versionNumber, commit, date)
}
