package test

import (
	"testing"

	"bookcollection.com/rest-api/mathservices"
)

// Simply here to learn Unit testing in Go. My course didn't cover this.
func TestAddition(t *testing.T) {
	a := 6
	b := 6
	want := mathservices.AddTwoIntegers(a, b)

	if want != 12 {
		t.Error("Expected 12, got", want)
	}
}
