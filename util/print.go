package util

import (
	"fmt"

	"github.com/fatih/color"
)

// PrintfOk formats according to a format specifier f and writes to
// color.Output with a prefix 'OK' in green color.
// It returns the number of bytes written and any write error encountered.
func PrintfOk(f string, a ...interface{}) (int, error) {
	return fprintfc("%s %s", color.GreenString("[OK]"), fmt.Sprintf(f, a...))
}

// PrintfInfo formats according to a format specifier f and writes to
// color.Output with a prefix 'INFO' in blue color.
// It returns the number of bytes written and any write error encountered.
func PrintfInfo(f string, a ...interface{}) (int, error) {
	return fprintfc("%s %s", color.BlueString("[INFO]"), fmt.Sprintf(f, a...))
}

// PrintfError formats according to a format specifier f and writes to
// color.Output with a prefix 'ERROR' in red color.
// It returns the number of bytes written and any write error encountered.
func PrintfError(f string, a ...interface{}) (int, error) {
	return fprintfc("%s %s", color.RedString("[ERROR]"), fmt.Sprintf(f, a...))
}

func fprintfc(f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(color.Output, f, a...)
}
