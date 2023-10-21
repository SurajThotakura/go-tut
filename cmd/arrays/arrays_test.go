package arrays

import "testing"

func TestArraySum(t *testing.T) {
	inputNumbersArray := [5]int{1, 2, 3, 4, 5}
	got := ArraySum(inputNumbersArray)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestSliceSum(t *testing.T) {
	inputSlice := []int{1, 2, 10}
	got := SliceSum(inputSlice)
	want := 13

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSliceOfSums(t *testing.T) {
	checkSum := func(sliceOne []int, sliceTwo []int) bool {
		for i := range sliceOne {
			if sliceOne[i] != sliceTwo[i] {
				return false
			}
		}
		return true
	}

	got := SliceOfSums([]int{1, 2, 10}, []int{43, 65, 10}, []int{1, 10})
	want := []int{13, 118, 11}

	if !checkSum(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
