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
	t.Run("rectangle area", func(t *testing.T){
		rect := Rectangle{12.0, 6.0}
		got := rect.Area()
		want := 72.0

		if want != got {
			t.Errorf("Wanted %.2f, got %.2f", want, got)
		}
	})

	t.Run("circle area", func(t *testing.T){
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if want != got {
			t.Errorf("Got %g, Want %g", got, want)
		}
	})
}
