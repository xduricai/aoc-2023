package util

type QueueNode[T any] struct {
	next  *QueueNode[T]
	value T
}

type Queue[T any] struct {
	head   *QueueNode[T]
	tail   *QueueNode[T]
	length int
}

func newQueueNode[T any](value *T) *QueueNode[T] {
	return &QueueNode[T]{
		next:  nil,
		value: *value,
	}
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (queue *Queue[T]) Enqueue(value *T) {
	node := newQueueNode[T](value)

	if queue.length == 0 {
		queue.head = node
	} else {
		queue.tail.next = node
	}
	queue.tail = node
	queue.length++
}

func (queue *Queue[T]) Dequeue() *T {
	if queue.length == 0 {
		return nil
	}
	if queue.length == 1 {
		queue.tail = nil
	}

	val := queue.head.value
	queue.head = queue.head.next
	queue.length--

	return &val
}

func (queue *Queue[T]) Length() int {
	return queue.length
}
