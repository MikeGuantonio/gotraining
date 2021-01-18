package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if want != got {
		t.Errorf("Wanted %.2f, Got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	if want != got {
		t.Errorf("Wanted %.2f, got %.2f", want, got)
	}
}