//go:build !release

package assure

import "fmt"

func Assert(condition func() bool, msg string) {
	if !condition() {
		failWithMsg(msg)
	}
}

func Equal(a any, b any, msg string) {
	if a != b {
		failWithMsg(msg)
	}
}

func NotEqual(a any, b any, msg string) {
	if a == b {
		failWithMsg(msg)
	}
}

func Nil(a any, msg string) {
	if a != nil {
		failWithMsg(msg)
	}
}

func NotNil(a any, msg string) {
	if a == nil {
		failWithMsg(msg)
	}
}

func Err(err error, msg string) {
	if err == nil {
		failWithMsg(msg)
	}
}

func NoErr(err error, msg string) {
	if err != nil {
		failWithMsg(msg)
	}
}

func failWithMsg(msg string) {
	panic(fmt.Sprintf("Assetion failed: %s", msg))
}
