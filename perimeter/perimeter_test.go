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

	areaTests := map[string]struct {
		shape Shape
		want float64
	}{
		"rectangle area": {
			shape: Rectangle{12, 6},
			want: 72.0,
		},
		"circle area": {	
			shape: Circle{10},
			want: 314.1592653589793,
		},
		"triangle area" : {
			shape: Triangle{12, 6},
			want: 36.0,
		},
	}

	for name, tt := range areaTests {
		t.Run(name, func(t *testing.T){
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%s: Wanted %.2f, got %.2f", name, tt.want, got)
			}
		})
	}
}
