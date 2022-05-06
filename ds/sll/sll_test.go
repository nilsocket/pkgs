package sll

import (
	"reflect"
	"testing"
)

func assertData(t testing.TB, s *SLL, want []int) {
	t.Helper()
	got := s.ToSlice()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("sll data: %v, want:%v", s, want)
	}

	if s.len != len(want) {
		t.Errorf("got Len: %v, want Len:%v", s.len, len(want))
	}
}

func TestNew(t *testing.T) {

	s := New()
	assertData(t, s, []int{})

	s = New(1, 2)
	assertData(t, s, []int{1, 2})
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
			s := New()
			s.Append(v.input...)

			assertData(t, s, v.output)
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

		s := New(1)

		for _, v := range tests {
			s.Append(v.input...)

			assertData(t, s, v.output)
		}
	})

	t.Run("Check for tail changes on append", func(t *testing.T) {
		s := New(1)
		pt := s.Tail()
		s.Append(2)
		ct := s.Tail()

		if pt == ct {
			t.Error("tail not changed on Append(")
		}
	})
}

func TestToSlice(t *testing.T) {
	s := New(1, 2, 3, 4, 5, 6, 7)
	want := []int{1, 2, 3, 4, 5, 6, 7}

	assertData(t, s, want)
}

// TestInsertAt also test for head and tail pointers
func TestInsertAt(t *testing.T) {

	t.Run("Passing nil to InsertAt()", func(t *testing.T) {
		s := New()
		s.InsertAfter(s.Head(), 1)
		want := []int{}

		assertData(t, s, want)
	})

	t.Run("Passing Tail to InsertAt()", func(t *testing.T) {
		s := New(1)
		n := s.InsertAfter(s.Tail(), 2, 3)
		want := []int{1, 2, 3}

		assertData(t, s, want)

		if n != s.Tail() {
			t.Error("tail not changed on InsertAt() operation")
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
			s := New()
			s.Prepend(v.input...)

			assertData(t, s, v.output)
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

		s := New(1)

		for _, v := range tests {
			s.Prepend(v.input...)

			assertData(t, s, v.output)
		}
	})

	t.Run("Check for head changes on prepend", func(t *testing.T) {
		s := New(1)
		ph := s.Head()
		s.Prepend(2, 3)
		ch := s.Head()
		if ph == ch {
			t.Error("head not changed on Prepend()")
		}
	})
}

func TestDeleteN(t *testing.T) {

	t.Run("Delete some elements", func(t *testing.T) {
		s := New(1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 5)
		s.Delete(1)
		s.DeleteN(2, 2)
		s.DeleteN(3, 2)
		s.Delete(5)

		want := []int{1, 2, 2, 3, 4}

		assertData(t, s, want)
	})

	t.Run("Delete all elements", func(t *testing.T) {
		s := New(1, 2)
		s.DeleteN(1, 1)
		s.DeleteN(2, 1)

		want := []int{}

		assertData(t, s, want)

		if s.head != nil && s.tail != nil {
			t.Error("head and tail were supposed to be nil")
		}
	})

}
