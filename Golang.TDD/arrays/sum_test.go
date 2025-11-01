package arrays

import (
	"reflect"
	"testing"
)

/*
Arrays have a fixed capacity which you define when you declare the variable.
We can initialize an array in two ways:
    [N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}

    [...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}
*/

func TestSum(t *testing.T) {

	t.Run("collections of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSummAllTail(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3, 4}, []int{0, 9, 8, 7})
		want := []int{9, 24}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
