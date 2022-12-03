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
	if expected != actual {
		a.t.Errorf("expected %d but got %d", expected, actual)
	}
}

func (a Assert) EqualInt(expected, actual int) {
	if expected != actual {
		a.t.Errorf("expected %d but got %d", expected, actual)
	}
}

func (a Assert) EqualString(expected, actual string) {
	if expected != actual {
		a.t.Errorf("expected %s but got %s", expected, actual)
	}
}

func (a Assert) True(cond bool) {
	if !cond {
		a.t.Error("expected true")
	}
}

func (a Assert) NotError(err error) {
	if err != nil {
		a.t.Error(err)
	}
}
