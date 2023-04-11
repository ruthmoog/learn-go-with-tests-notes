package structsmethodsandinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("wanted %.2f, but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expectedArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expectedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expectedArea {
				t.Errorf("%#v wanted area %.2f, but got %.2f", tt.shape, tt.expectedArea, got)
			}
		})
	}
}
