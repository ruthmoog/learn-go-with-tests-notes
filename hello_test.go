package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ruth")
	want := "Hello, Ruth!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
