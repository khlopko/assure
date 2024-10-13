//go:build release

package assure

func Assert(condition func() bool, msg string) {
}

func Equal(a any, b any, msg string) {
}

func NotEqual(a any, b any, msg string) {
}

func Nil(a any, msg string) {
}

func NotNil(a any, msg string) {
}

func Err(err error, msg string) {
}

func NoErr(err error, msg string) {
}

