package sll

import (
	"bytes"
	"fmt"
	"reflect"
)

type Element struct {
	Value int
	next  *Element
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) String() string {
	return fmt.Sprint(e.Value)
}

type SLL struct {
	head *Element
	tail *Element
	len  int
}

func New(vs ...int) *SLL {
	var s SLL
	s.Append(vs...)
	return &s
}

func (s *SLL) Head() *Element {
	return s.head
}

func (s *SLL) Tail() *Element {
	return s.tail
}

func (s *SLL) Len() int {
	return s.len
}

func (s *SLL) firstElement(v int) *Element {
	ne := &Element{Value: v}
	s.head = ne
	s.tail = ne
	s.len++

	return ne
}

// Append v to sll
// returns last inserted element
func (s *SLL) Append(vs ...int) *Element {
	if len(vs) == 0 {
		return nil
	}

	if s.head == nil || s.tail == nil {
		s.firstElement(vs[0])
		if len(vs) > 1 {
			return s.InsertAfter(s.head, vs[1:]...)
		}
		return s.head
	}
	return s.InsertAfter(s.tail, vs...)
}

// InsertAfter, insert vs elements at e; e can't be nil
// returns last inserted element
func (s *SLL) InsertAfter(e *Element, vs ...int) *Element {
	if e == nil {
		return nil
	}
	return s.insertAfter(e, vs...)
}

func (s *SLL) insertAfter(e *Element, vs ...int) *Element {

	oe := e

	for i := range vs {
		e = insertAfter(e, vs[i])
	}

	if s.tail == oe {
		s.tail = e
	}
	s.len += len(vs)

	return e
}

func insertAfter(e *Element, v int) *Element {
	ne := &Element{Value: v}
	ne.next = e.next
	e.next = ne
	return ne
}

// Prepend, prepend `vs` elements to `s.head`
func (s *SLL) Prepend(vs ...int) *Element {
	if len(vs) == 0 {
		return nil
	}

	ts := New(vs...)
	le := ts.tail
	ts = CombineLists(ts, s)

	s.head = ts.head
	s.tail = ts.tail
	s.len = ts.len
	return le
}

// Delete, one  element with value `v`
func (s *SLL) Delete(v int) {
	s.DeleteN(v, 1)
}

// DeleteN, delete `n` elements with value `v`, `n` times
// if `n` == -1, all repetitions of `v` are deleted
func (s *SLL) DeleteN(v int, n int) {
	var pe *Element
	for ce := s.head; n != 0 && ce != nil; ce = ce.next {
		if reflect.DeepEqual(ce.Value, v) {
			s.delete(pe, ce)
			n--
		} else {
			pe = ce
		}
	}
}

func (s *SLL) delete(pe, ce *Element) {
	if pe == nil { // head
		s.head = s.head.next
	} else {
		pe.next = ce.next
	}
	if s.tail == ce {
		s.tail = pe
	}
	s.len--
}

func CombineLists(a, b *SLL) *SLL {
	if a.head == nil || a.tail == nil {
		return b
	} else if b.head == nil || b.tail == nil {
		return a
	}

	a.tail.next = b.head
	a.tail = b.tail
	a.len += b.len
	return a
}

func (s *SLL) ToSlice() []int {
	res := make([]int, 0, s.len)
	for t := s.head; t != nil; t = t.next {
		res = append(res, t.Value)
	}
	return res
}

func (s *SLL) String() string {
	buf := &bytes.Buffer{}

	fmt.Fprint(buf, "[")
	for t, i := s.head, s.len; i > 0; t, i = t.next, i-1 {
		fmt.Fprint(buf, t, " ")
	}
	fmt.Fprint(buf, s.tail, "]")

	return buf.String()
}
