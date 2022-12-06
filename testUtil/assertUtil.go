package testUtil

import "testing"

func AssertEqual[T comparable](t *testing.T, expected T, actual T) {
	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
