package heap

import (
	"reflect"
	"sort"
	"testing"

	"github.com/nilsocket/pkgs/utils/random"
)

const (
	maxNum = 1e6
	nums   = 1e6
)

var data = random.Intsn(nums, maxNum)

func TestHeapSort(t *testing.T) {
	sortedInts := makeCopy(data)
	sort.Ints(sortedInts)

	testHeapSort(t, data, sortedInts)
}

func testHeapSort(t *testing.T, input, want []int) {
	t.Helper()

	inputCopy := makeCopy(input)

	MaxSort(input)
	got := input

	assertEqual(t, inputCopy, got, want)
}

func assertEqual(t *testing.T, input, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("input: %v\ngot : %v\nwant: %v", input, got, want)
	}
}

func makeCopy(x []int) []int {
	t := make([]int, len(x))
	copy(t, x)
	return t
}

// benchmark

func BenchmarkMaxSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := random.Ints(nums)
		b.StartTimer()
		MaxSort(data)
	}
}
func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := random.Ints(nums)
		b.StartTimer()
		sort.Ints(data)
	}
}
