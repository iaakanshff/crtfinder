package runner

import (
	"fmt"
	"os"
)

// Checking version
func CheckVersion() {
	version := fmt.Sprintf(Succfix + "Current crtfinder version: " + g.BrGreen("v1.0"))
	fmt.Fprintln(os.Stderr, version)
}
