package stack_test

import (
	"reflect"
	"testing"

	"github.com/nilsocket/pkgs/ds/stack"
)

func TestPushPOP(t *testing.T) {
	inp := []any{1, 2, 3, 4, 5, 6, 7}
	want := []any{7, 6, 5, 4, 3, 2, 1}
	var got []any

	s := stack.New()

	for _, v := range inp {
		s.Push(v)
	}

	for v, err := s.Pop(); err == nil; v, err = s.Pop() {
		got = append(got, v)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
