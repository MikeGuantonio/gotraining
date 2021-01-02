package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("Want %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}