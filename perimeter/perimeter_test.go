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

	areaTests := []struct {
		description string
		shape Shape
		want float64
	}{
		{
			"rectangle area",
			Rectangle{12, 6},
			72.0,
		},
        {
			"circle area",
			Circle{10},
			314.1592653589793,
		},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Wanted %.2f, got %.2f", tt.want, got)
		}
	}
}
