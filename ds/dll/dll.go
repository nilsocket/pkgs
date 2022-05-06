package dll

import (
	"bytes"
	"fmt"
	"reflect"
)

type Element struct {
	prev  *Element
	Value int
	next  *Element
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

func (e *Element) String() string {
	return fmt.Sprint(e.Value)
}

type DLL struct {
	head *Element
	tail *Element
	len  int
}

func New(vs ...int) *DLL {
	var s DLL
	s.Append(vs...)
	return &s
}

func (d *DLL) Head() *Element {
	return d.head
}

func (d *DLL) Tail() *Element {
	return d.tail
}

func (d *DLL) Len() int {
	return d.len
}

func (d *DLL) firstElement(v int) *Element {
	ne := &Element{Value: v}
	d.head = ne
	d.tail = ne
	d.len++

	return ne
}

// Append v to sll
// returns last inserted element
func (d *DLL) Append(vs ...int) *Element {
	if len(vs) == 0 {
		return nil
	}

	if d.head == nil || d.tail == nil {
		d.firstElement(vs[0])
		if len(vs) > 1 {
			return d.InsertAfter(d.head, vs[1:]...)
		}
		return d.head
	}
	return d.InsertAfter(d.tail, vs...)
}

// InsertAfter, insert vs elements at e; e can't be nil
// returns last inserted element
func (d *DLL) InsertAfter(e *Element, vs ...int) *Element {
	if e == nil {
		return nil
	}
	return d.insertAfter(e, vs...)
}

func (d *DLL) insertAfter(e *Element, vs ...int) *Element {
	oe := e

	for i := range vs {
		e = insertAfter(e, vs[i])
	}

	if d.tail == oe {
		d.tail = e
	}
	d.len += len(vs)

	return e
}

func insertAfter(e *Element, v int) *Element {
	ne := &Element{Value: v}
	ne.next = e.next
	ne.prev = e
	e.next = ne

	if ne.next != nil {
		ne.next.prev = ne
	}

	return ne
}

func (d *DLL) InsertBefore(e *Element, vs ...int) *Element {
	if e == nil {
		return nil
	}
	return d.insertBefore(e, vs...)
}

func (d *DLL) insertBefore(e *Element, vs ...int) *Element {
	oe := e
	for i := range vs {
		e = insertBefore(e, vs[i])
	}
	if d.head == oe {
		d.head = e
	}
	d.len += len(vs)

	return e
}

func insertBefore(e *Element, v int) *Element {
	ne := &Element{Value: v}
	ne.next = e
	ne.prev = e.prev
	e.prev = ne

	if ne.prev != nil {
		ne.prev.next = ne
	}

	return ne
}

// Prepend, prepend `vs` elements to `s.head`
func (d *DLL) Prepend(vs ...int) *Element {
	if len(vs) == 0 {
		return nil
	}

	ts := New(vs...)
	le := ts.tail
	ts = CombineLists(ts, d)

	d.head = ts.head
	d.tail = ts.tail
	d.len = ts.len
	return le
}

// Delete, one  element with value `v`
func (s *DLL) Delete(v int) {
	s.DeleteN(v, 1)
}

// DeleteN, delete `n` elements with value `v`, `n` times
// if `n` == -1, all repetitions of `v` are deleted
func (d *DLL) DeleteN(v int, n int) {
	var pe *Element
	for ce := d.head; n != 0 && ce != nil; ce = ce.next {
		if reflect.DeepEqual(ce.Value, v) {
			d.delete(pe, ce)
			n--
		} else {
			pe = ce
		}
	}
}

func (d *DLL) delete(pe, ce *Element) {
	if pe == nil { // head
		d.head = d.head.next
	} else {
		pe.next = ce.next
	}
	if d.tail == ce {
		d.tail = pe
	}
	d.len--
}

func CombineLists(a, b *DLL) *DLL {
	if a.head == nil || a.tail == nil {
		return b
	} else if b.head == nil || b.tail == nil {
		return a
	}

	a.tail.next = b.head
	b.head.prev = a.tail

	a.tail = b.tail
	a.len += b.len
	return a
}

func (d *DLL) ToSlice() []int {
	res := make([]int, 0, d.len)
	for t := d.head; t != nil; t = t.next {
		res = append(res, t.Value)
	}
	return res
}

func (d *DLL) ToSliceReverse() []int {
	res := make([]int, 0, d.len)
	for t := d.tail; t != nil; t = t.prev {
		res = append(res, t.Value)
	}
	return res
}

func (d *DLL) String() string {
	buf := &bytes.Buffer{}

	fmt.Fprint(buf, "[")
	for t, i := d.head, d.len; i > 1; t, i = t.next, i-1 {
		fmt.Fprint(buf, t, " ")
	}
	fmt.Fprint(buf, d.tail, "]")

	return buf.String()
}
