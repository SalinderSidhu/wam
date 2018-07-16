package util

import (
	"fmt"

	"github.com/fatih/color"
)

// PrintOk
func PrintOk(f string, a ...interface{}) (int, error) {
	return fprintc("%s %s", color.GreenString("[OK]"), fmt.Sprintf(f, a...))
}

// PrintInfo
func PrintInfo(f string, a ...interface{}) (int, error) {
	return w.fprintc("%s %s", color.BlueString("[INFO]"), fmt.Sprintf(f, a...))
}

// PrintError
func PrintError(f string, a ...interface{}) (int, error) {
	return w.fprintc("%s %s", color.RedString("[ERROR]"), fmt.Sprintf(f, a...))
}

func fprintc(f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(color.Output, f, a...)
}
