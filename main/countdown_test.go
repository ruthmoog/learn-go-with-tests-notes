package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("Not enough calls to sleeper, wanted 3 but got %q", spySleeper.Calls)
	}
}
