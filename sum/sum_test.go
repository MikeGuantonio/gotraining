package sum

import "testing"

func TestSum(t *testing.T) {
	assertIsEqual := func(t *testing.T, got int, want int, numbers []int) {
		t.Helper()
		if want != got {
			t.Errorf("Wanted %d got %d given %v", want, got, numbers)
		}
	}

	t.Run("Should sum a collection of any size", func(t *testing.T){
		numbers := []int{1,2,3}

		got := Sum(numbers)
		want := 6

		assertIsEqual(t, got, want, numbers)
	})
	
}