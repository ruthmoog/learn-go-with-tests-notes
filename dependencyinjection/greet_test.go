package dependencyinjection

//package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ruth")

	got := buffer.String()
	want := "Hi Ruth"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}
