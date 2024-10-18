package assure_test

import (
	"errors"
	"testing"

	"github.com/khlopko/assure"
)

func Recover(f func(), t *testing.T) {
	t.Helper()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Assertion expected to panic")
		}
	}()

	f()
}

func TestAssertPanics_WhenFalse(t *testing.T) {
	Recover(func() {
		assure.Assert(func() bool { return false }, "")
	}, t)
}

func TestEqualPanics_WhenNotEqual(t *testing.T) {
	Recover(func() {
		assure.Equal(1, 2, "")
		assure.Equal("a", "b", "")
		assure.Equal(1, "b", "")
	}, t)
}

func TestNotEqualPanics_WhenEqual(t *testing.T) {
	Recover(func() {
		assure.NotEqual(1, 1, "")
		assure.NotEqual("a", "a", "")
	}, t)
}

func TestNilPanics_WhenNotNil(t *testing.T) {
	Recover(func() {
		assure.Nil(1, "")
	}, t)
}

func TestNotNilPanics_WhenNil(t *testing.T) {
	Recover(func() {
		assure.NotNil(nil, "")
	}, t)
}

func TestErrPanics_WhenNoErr(t *testing.T) {
	Recover(func() {
		assure.Err(nil, "")
	}, t)
}

func TestNoErrPanics_WhenErr(t *testing.T) {
	Recover(func() {
		assure.NoErr(errors.New("my error"), "")
	}, t)
}

func TestNeverPanics(t *testing.T) {
	Recover(func() {
		assure.Never("you got it all wrong")
	}, t)
}

func TestDisabledDefaultContext(t *testing.T) {
	assure.SetIsEnabled(false)
	defer assure.SetIsEnabled(true)

	assure.Assert(func() bool { return false }, "")
	assure.Equal(1, 2, "")
	assure.NotEqual(1, 1, "")
	assure.Nil(1, "")
	assure.NotNil(nil, "")
	assure.Err(nil, "")
	assure.NoErr(errors.New("my error"), "")
	assure.Never("really never")
}

func TestCustomDefaultFailHandler(t *testing.T) {
	handled := false
	assure.SetFailHandler(func(msg string) {
		handled = true
	})

	assure.Assert(func() bool { return false }, "")

	if !handled {
		t.Error("Expect to change `handled` variable to `true` in a custom handler")
	}
}
