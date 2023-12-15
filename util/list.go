package util

type Node[T comparable] struct {
	Next  *Node[T]
	Value T
}

type LinkedList[T comparable] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		Next:  nil,
		Value: value,
	}
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (list *LinkedList[T]) Push(value T) {
	node := NewNode[T](value)

	if list.Length == 0 {
		list.Head = node
	} else {
		list.Tail.Next = node
	}

	list.Tail = node
	list.Length++
}

func (list *LinkedList[T]) Remove(value T) {
	if list.Length == 0 {
		return
	}

	if list.Head.Value == value {
		if list.Length == 1 {
			list.Tail = nil
		}
		list.Head = list.Head.Next
		list.Length--
		return
	}

	previous := list.Head
	current := list.Head.Next

	for current != nil {
		if current.Value == value {
			previous.Next = current.Next
			list.Length--

			if current == list.Tail {
				list.Tail = previous
			}
			return
		}
		previous = current
		current = current.Next
	}
}
