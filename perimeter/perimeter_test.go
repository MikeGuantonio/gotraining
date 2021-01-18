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
	rect := Rectangle{12.0, 6.0}
	got := Area(rect)
	want := 72.0

	if want != got {
		t.Errorf("Wanted %.2f, got %.2f", want, got)
	}
}