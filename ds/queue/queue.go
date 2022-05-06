package queue

type Queue struct {
	data []any
	head int
	len  int
}

func New(cap int) *Queue {
	return &Queue{data: make([]any, cap)}
}

func (q *Queue) Cap() int {
	return cap(q.data)
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Peek() any {
	if q.len <= 0 {
		return nil
	}
	return q.data[q.head]
}

func (q *Queue) Pop() any {
	if q.len <= 0 {
		return nil
	}
	q.len--
	q.head = (q.head + 1) % cap(q.data)
	return q.data[q.head]
}

func (q *Queue) Push(d any) {
	q.len++
	if q.len >= len(q.data) {
		tbuf := make([]any, q.len*2)
		copy(tbuf, q.data[q.head:])
		copy(tbuf, q.data[:q.head])
		q.data = tbuf
		q.head = 0
	}
	q.data[(q.head+q.len)%len(q.data)] = d
}
