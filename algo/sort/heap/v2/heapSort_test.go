package heap

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/nilsocket/pkgs/ds/heap/v2"
)

const (
	maxNum = 1e4
	nums   = 1e4
)

var data = generateData(nums, maxNum)

func TestSort(t *testing.T) {
}

func TestHeapSort(t *testing.T) {
	sortedInts := makeCopy(data)
	sort.Sort(tt(sortedInts))

	testHeapSort(t, data, sortedInts)
}

type tt []heap.Elem[int]

func (t tt) Less(i, j int) bool { return t[i].Key < t[j].Key }
func (t tt) Len() int           { return len(t) }
func (t tt) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func testHeapSort(t *testing.T, input, want []heap.Elem[int]) {
	t.Helper()

	inputCopy := makeCopy(input)

	Sort(input, heap.MaxHeap)
	got := input

	assertEqual(t, inputCopy, got, want)
}

func assertEqual(t *testing.T, input, got, want []heap.Elem[int]) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("input: %v\ngot : %v\nwant: %v", input, got, want)
	}
}

func makeCopy(x []heap.Elem[int]) []heap.Elem[int] {
	t := make([]heap.Elem[int], len(x))
	copy(t, x)
	return t
}

func generateData(size, n int) []heap.Elem[int] {
	res := make([]heap.Elem[int], size)
	for size > 0 {
		size--
		res[size] = heap.Elem[int]{Key: rand.Intn(n), Data: nil}
	}
	return res
}

// benchmark

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := generateData(nums, maxNum)
		b.StartTimer()
		Sort(data, heap.MaxHeap)
	}
}
func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := generateData(nums, maxNum)
		b.StartTimer()
		sort.Sort(tt(data))
	}
}
