package binary

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	input := []int{1, 3, 5, 7, 9, 10, 11}

	t.Run("Searching for Existing Elements:", func(t *testing.T) {
		searchElems := []int{1, 7, 9, 11} // start, mid, other, end
		want := []int{0, 3, 4, 6}
		assertData(t, input, searchElems, want)
	})

	t.Run("Searching for Non-Existing Elements:", func(t *testing.T) {
		searchElems := []int{0, 3, 2, 4, 10, 12}
		want := []int{-1, 1, -1, -1, 5, -1}
		assertData(t, input, searchElems, want)
	})
}

func assertData(t testing.TB, input, searchElems, want []int) {
	t.Helper()

	got := multiSearch(input, searchElems)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v", got, want)
	}
}

func multiSearch(a []int, elems []int) []int {
	got := make([]int, len(elems))
	for i := range elems {
		got[i] = Search(a, elems[i])
	}
	return got
}
