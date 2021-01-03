package sum

import "testing"

func TestSum(t *testing.T) {
	assertIsEqual := func(t *testing.T, got int, want int, numbers [5]int) {
		t.Helper()
		if want != got {
			t.Errorf("Wanted %d got %d given %v", want, got, numbers)
		}
	}
	t.Run("Sum a list of integers", func(t *testing.T){
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertIsEqual(t, got, want, numbers)
	})
	
}