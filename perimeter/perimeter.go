package perimeter

const PI = 3.14

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rect Rectangle) float64{
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}

func Area(circle Circle) float64 {
	return circle.Radius * PI
}