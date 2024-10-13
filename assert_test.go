package assure_test

import (
	"testing"
	"assure"
)

func TestAssertPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Assertion expected to panic")
		}
	}()

	assure.Assert(false, "i should have fail")
}
