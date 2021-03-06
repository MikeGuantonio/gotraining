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

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
	}
	
	t.Run("Should sum all tail values in an array", func(t *testing.T){
		slice1 := []int {1, 2}		
		slice2 := []int {2, 9}

		got := SumAllTails(slice1, slice2)
		want := []int {2, 9}

		checkSums(t, got, want)
	})

	t.Run("Should sum all tails for slices of variable length", func(t *testing.T){
		slice1 := []int {1, 2, 4}
		slice2 := []int {5, 8, 10, 12}

		got := SumAllTails(slice1, slice2)
		want := []int {6, 30}

		checkSums(t, got, want)
	})

	t.Run("Should sum an empty slice", func(t *testing.T) {
		slice1 := []int {}

		got := SumAllTails(slice1)
		want := []int {0}

		checkSums(t, got, want)
	})
}