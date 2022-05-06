package heap

import (
	"reflect"
	"sort"
	"testing"

	"github.com/nilsocket/pkgs/utils/random"
)

const (
	maxNum = 1e6
	nums   = 1e4
)

var data = sort.IntSlice(random.Intsn(nums, maxNum))

func TestHeapSort(t *testing.T) {
	sortedInts := makeCopy(data)
	sort.Ints(sortedInts)
	reverse(sortedInts)
	testHeapSort(t, data, sortedInts)
}

func reverse(res []int) {
	for i := 0; i < len(res)/2; i++ {
		r := len(res) - 1 - i
		res[i], res[r] = res[r], res[i]
	}
}

func testHeapSort(t *testing.T, input, want []int) {
	t.Helper()

	inputCopy := makeCopy(input)

	Sort(sort.IntSlice(input))
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

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := random.Ints(nums)
		b.StartTimer()
		Sort(sort.IntSlice(data))
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
