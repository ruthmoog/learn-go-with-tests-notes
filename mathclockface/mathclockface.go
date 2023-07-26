package mathclockface

import (
	"math"
	"time"
)

const secondsInHalfClock = 30
const secondsInClock = secondsInHalfClock * 2
const minutesInHalfClock = 30
const minutesInClock = minutesInHalfClock * 2
const hoursInHalfClock = 6
const hoursInClock = hoursInHalfClock * 2

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue
// clock at time `t` represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // Scale
	p = Point{p.X, -p.Y}                                      // Flip horizontal for SVG top left origin
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // Translate position
	return p
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second())) // One second is (2π / 60) radians
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
