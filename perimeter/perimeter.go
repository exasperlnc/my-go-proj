package perimeter 

// import "math"

// Perimiter takes the width and height of a rectangle and returns the perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area takes the width and height of a rectangle and returns the area
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

type Rectangle struct {
	Width float64
	Height float64
}