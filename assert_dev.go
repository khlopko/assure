//go:build !release
// +build !release

package assure

import "fmt"

func Assert(condition bool, msg string) {
	if !condition {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

