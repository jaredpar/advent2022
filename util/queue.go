package util

type queueNode[T any] struct {
	value T
	next  *queueNode[T]
}

func newQueueNode[T any](value T, next *queueNode[T]) *queueNode[T] {
	return &queueNode[T]{value: value, next: next}
}

type Queue[T any] struct {
	head *queueNode[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{head: nil}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue[T]) Enqueue(value T) {
	q.head = newQueueNode(value, q.head)
}

func (q *Queue[T]) Dequeue() T {
	if q.head == nil {
		panic("can't dequeue an empty queue")
	}

	cur := q.head
	q.head = cur.next
	return cur.value
}
