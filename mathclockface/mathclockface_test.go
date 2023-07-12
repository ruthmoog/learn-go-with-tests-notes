package mathclockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)

	want := math.Pi
	got := secondsInRadians(thirtySeconds)

	if got != want {
		t.Fatalf("Wanted %v radians, but got %v", want, got)
	}
}
