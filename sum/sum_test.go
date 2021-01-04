package sum

import (
	"testing"
	"reflect"
)

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

func TestSumAll(t *testing.T) {
	t.Run("Should sum each slice and place the result in a single list", func(t *testing.T){
		slice1 := []int {1, 2}
		slice2 := []int {0, 9}

		got := SumAll(slice1, slice2)
		want := []int {3, 9}
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted %d, Got %d, Supplied %v...", got, want, slice1)
		}
	})
}