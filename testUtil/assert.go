package testUtil

import (
	"testing"
)

type Assert struct {
	t *testing.T
}

func NewAssert(t *testing.T) Assert {
	return Assert{t: t}
}

func (a Assert) EqualByte(expected, actual byte) {
	AssertEqual(a.t, expected, actual)
}

func (a Assert) EqualInt(expected, actual int) {
	AssertEqual(a.t, expected, actual)
}

func (a Assert) EqualString(expected, actual string) {
	AssertEqual(a.t, expected, actual)
}

func (a Assert) True(cond bool) {
	if !cond {
		a.t.Error("expected true")
	}
}

func (a Assert) NotError(err error) {
	if err != nil {
		a.t.Fatal(err)
	}
}
