package debug

import "fmt"

// `fmt.Printf`
func Printf(format string, args ...interface{}) {
	if debug {
		fmt.Printf(format, args...)
	}
}
