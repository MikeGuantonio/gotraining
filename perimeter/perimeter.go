package perimeter

type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(rect Rectangle) float64{
	return 2 * (rect.Width + rect.Height)
}

func Area(width float64, height float64) float64 {
	return width * height
}