package assure_test

import (
	"assure"
	"errors"
	"testing"
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
