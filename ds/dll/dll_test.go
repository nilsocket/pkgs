package dll_test

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"

	"github.com/nilsocket/pkgs/ds/dll"
	"github.com/nilsocket/pkgs/utils/random"
)

func assertData(t testing.TB, d *dll.DLL, want []int) {
	t.Helper()
	got := d.ToSlice()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("sll data: %v, want:%v", got, want)
	}

	if d.Len() != len(want) {
		t.Errorf("got Len: %v, want Len:%v", d.Len(), len(want))
	}
}

func TestNew(t *testing.T) {

	d := dll.New()
	assertData(t, d, []int{})

	d = dll.New(1, 2)
	assertData(t, d, []int{1, 2})
}

func TestAppend(t *testing.T) {

	t.Run("Append with empty SLL's", func(t *testing.T) {
		tests := []struct {
			input, output []int
		}{
			{input: []int{}, output: []int{}},
			{input: []int{1, 2, 3, 4}, output: []int{1, 2, 3, 4}},
		}

		for _, v := range tests {
			d := dll.New()
			d.Append(v.input...)

			assertData(t, d, v.output)
		}
	})

	t.Run("Append with non-empty SLL's", func(t *testing.T) {
		tests := []struct {
			input, output []int
		}{
			{input: []int{}, output: []int{1}},
			{input: []int{2, 3, 4}, output: []int{1, 2, 3, 4}},
			{input: []int{5}, output: []int{1, 2, 3, 4, 5}},
			{input: []int{6, 7}, output: []int{1, 2, 3, 4, 5, 6, 7}},
		}

		d := dll.New(1)

		for _, v := range tests {
			d.Append(v.input...)

			assertData(t, d, v.output)
		}
	})

	t.Run("Check for tail changes on append", func(t *testing.T) {
		d := dll.New(1)
		pt := d.Tail()
		d.Append(2)
		ct := d.Tail()

		if pt == ct {
			t.Error("tail not changed on Append(")
		}
	})
}

func TestToSlice(t *testing.T) {
	d := dll.New(1, 2, 3, 4, 5, 6, 7)
	want := []int{1, 2, 3, 4, 5, 6, 7}

	assertData(t, d, want)
}

func TestToSliceReverse(t *testing.T) {
	d := dll.New(1, 2, 3, 4, 5, 6, 7)
	want := []int{7, 6, 5, 4, 3, 2, 1}
	got := d.ToSliceReverse()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

// TestInsertAt also test for head and tail pointers
func TestInsertAfter(t *testing.T) {

	t.Run("Passing nil to InsertAfter()", func(t *testing.T) {
		d := dll.New()
		d.InsertAfter(d.Head(), 1)
		want := []int{}

		assertData(t, d, want)
	})

	t.Run("Passing Tail to InsertAfter()", func(t *testing.T) {
		d := dll.New(1)
		n := d.InsertAfter(d.Tail(), 2, 3)
		want := []int{1, 2, 3}

		assertData(t, d, want)

		if n != d.Tail() {
			t.Error("tail not changed on InsertAfter() operation")
		}
	})
}

func TestInsertBefore(t *testing.T) {

	t.Run("Passing nil to InsertBefore()", func(t *testing.T) {
		d := dll.New()
		d.InsertBefore(d.Head(), 1)
		want := []int{}

		assertData(t, d, want)
	})

	t.Run("Passing Head to InsertBefore()", func(t *testing.T) {
		d := dll.New(1)
		n := d.InsertBefore(d.Head(), 2, 3)
		want := []int{3, 2, 1}

		assertData(t, d, want)

		if n != d.Head() {
			t.Error("head not changed on InsertBefore() operation")
		}
	})
}

func TestPrepend(t *testing.T) {
	t.Run("Prepend with empty SLL's", func(t *testing.T) {
		tests := []struct {
			input, output []int
		}{
			{input: []int{}, output: []int{}},
			{input: []int{1, 2, 3, 4}, output: []int{1, 2, 3, 4}},
		}

		for _, v := range tests {
			d := dll.New()
			d.Prepend(v.input...)

			assertData(t, d, v.output)
		}
	})

	t.Run("Prepend with non-empty SLL's", func(t *testing.T) {
		tests := []struct {
			input, output []int
		}{
			{input: []int{}, output: []int{1}},
			{input: []int{2, 3, 4}, output: []int{2, 3, 4, 1}},
			{input: []int{5}, output: []int{5, 2, 3, 4, 1}},
			{input: []int{6, 7}, output: []int{6, 7, 5, 2, 3, 4, 1}},
		}

		d := dll.New(1)

		for _, v := range tests {
			d.Prepend(v.input...)

			assertData(t, d, v.output)
		}
	})

	t.Run("Check for head changes on prepend", func(t *testing.T) {
		d := dll.New(1)
		ph := d.Head()
		d.Prepend(2, 3)
		ch := d.Head()
		if ph == ch {
			t.Error("head not changed on Prepend()")
		}
	})
}

func TestDeleteN(t *testing.T) {

	t.Run("Delete some elements", func(t *testing.T) {
		d := dll.New(1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 5)
		d.Delete(1)
		d.DeleteN(2, 2)
		d.DeleteN(3, 2)
		d.Delete(5)

		want := []int{1, 2, 2, 3, 4}

		assertData(t, d, want)
	})

	t.Run("Delete all elements", func(t *testing.T) {
		d := dll.New(1, 2)
		d.DeleteN(1, 1)
		d.DeleteN(2, 1)

		want := []int{}

		assertData(t, d, want)

		if d.Head() != nil && d.Tail() != nil {
			t.Error("head and tail were supposed to be nil")
		}
	})
}

func TestString(t *testing.T) {
	d := dll.New(1, 2, 3, 4, 5, 6, 7)
	got := fmt.Sprint(d)
	want := "[1 2 3 4 5 6 7]"
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

/*
 * Benchmark
 */

var inp = random.Ints(1000000)

func BenchmarkAppend(b *testing.B) {
	d := dll.New()

	for i := 0; i < b.N; i++ {
		d.Append(inp...)
	}
}

func BenchmarkPushBack(b *testing.B) {
	l := list.New()

	for i := 0; i < b.N; i++ {
		for j := range inp {
			l.PushBack(inp[j])
		}
	}
}

func BenchmarkPrepend(b *testing.B) {
	d := dll.New()

	for i := 0; i < b.N; i++ {
		d.Prepend(inp...)
	}
}

func BenchmarkPushFront(b *testing.B) {
	l := list.New()

	for i := 0; i < b.N; i++ {
		for j := range inp {
			l.PushFront(inp[j])
		}
	}
}
