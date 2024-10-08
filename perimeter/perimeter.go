package perimeter 

import "math"

// Perimiter takes the width and height of a rectangle and returns the perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area takes the width and height of a rectangle and returns the area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}