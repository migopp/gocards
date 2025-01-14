package debug

import (
	"log"
	"os"
)

func Assert(c bool, m string) {
	if !c {
		log.Printf("Assertion failed: %s\n", m)
		os.Exit(1)
	}
}
