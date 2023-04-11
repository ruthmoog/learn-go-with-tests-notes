package structsmethodsandinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
