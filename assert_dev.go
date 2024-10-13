//go:build !release
// +build !release

package assure

import "fmt"

func Assert(condition func() bool, msg string) {
	if !condition() {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

func Equal(a any, b any, msg string) {
	if a != b {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

func NotEqual(a any, b any, msg string) {
	if a == b {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

func Nil(a any, msg string) {
	if a != nil {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

func NotNil(a any, msg string) {
	if a == nil {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

func Err(err error, msg string) {
	if err == nil {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

func NoErr(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

