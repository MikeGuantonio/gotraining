package iteration

import "testing"

func TestIteration(t *testing.T) {
	assertSame := func(t *testing.T, expected string, repeated string) {
		if expected != repeated {
			t.Errorf("Want %q got %q", expected, repeated)
		}
	}

	t.Run("Create a simple iteration", func (t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertSame(t, expected, repeated)
	})

	t.Run("Account for runs at zero", func(t *testing.T) {
		repeated := Repeat("b", 0)
		expected := ""
		assertSame(t, expected, repeated)
	})

	t.Run("Account for values less than zero", func(t *testing.T) {
		repeated := Repeat("c", -1)
		expected := ""
		assertSame(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}