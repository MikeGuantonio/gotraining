package perimeter

import (
	"math"
)

const PI = 3.14

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func(r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rect Rectangle) float64{
	return 2 * (rect.Width + rect.Height)
}