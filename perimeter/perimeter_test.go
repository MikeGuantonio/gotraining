package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if want != got {
		t.Errorf("Wanted %.2f, Got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if want != got {
			t.Errorf("Wanted %.2f, got %.2f", want, got)
		}
	}

	t.Run("rectangle area", func(t *testing.T){
		rect := Rectangle{12.0, 6.0}
		checkArea(t, rect, 72.0)
	})

	t.Run("circle area", func(t *testing.T){
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
