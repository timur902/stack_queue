package stack

type QueueInt struct {
	data[]int
}

func (q *QueueInt) Enqueue(x int) {
	q.data = append(q.data, x)
}

func (q *QueueInt) Len() int {
	return len(q.data)
}

func (q *QueueInt) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *QueueInt) Dequeue() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, true
}

func (q *QueueInt) Peek() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	return q.data[0], true
}

