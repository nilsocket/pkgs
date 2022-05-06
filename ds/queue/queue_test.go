package queue_test

import (
	"testing"

	"github.com/nilsocket/pkgs/ds/queue"
)

func TestQueue(t *testing.T) {

	t.Run("Push More data than Capacity", func(t *testing.T) {
		inp := []int{0, 1, 2, 3, 4, 5, 6}

		q := queue.New(5)
		push(q, inp)
		assertEqual(t, q, inp)
	})

	t.Run("Push and Pop data", func(t *testing.T) {
		inp1 := []int{0, 1, 2, 3, 4}
		inp2 := []int{5, 6, 7, 8, 9}
		want := []int{2, 3, 4, 5, 6, 7, 8, 9}

		q := queue.New(5)
		push(q, inp1)

		q.Pop()
		q.Pop()

		push(q, inp2)
		assertEqual(t, q, want)
	})

}

func push(q *queue.Queue, inp []int) {
	for i := range inp {
		q.Push(inp[i])
	}
}

func assertEqual(t *testing.T, q *queue.Queue, want []int) {
	t.Helper()
	if q.Len() != len(want) {
		t.Errorf("got len:%d, want len:%d", q.Len(), len(want))
	}
	for i := range want {
		if want[i] != q.Pop() {
			t.Errorf("got: %d, want:%d", q.Pop(), want[i])
		}
	}
}
