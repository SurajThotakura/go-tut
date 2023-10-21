package structs

import "math"

func Perimeter(rectangle Rectangle) float64 {
	perimeter := 2 * (rectangle.Height + rectangle.Width)
	return perimeter
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	area := r.Height * r.Width
	return area
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	area := math.Pi * c.Radius * c.Radius
	return area
}
