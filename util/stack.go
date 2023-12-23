package util

type StackNode[T any] struct {
	value T
	next  *StackNode[T]
}

func newStackNode[T any](value T) *StackNode[T] {
	return &StackNode[T]{
		value: value,
		next:  nil,
	}
}

type Stack[T any] struct {
	top    *StackNode[T]
	length int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		top:    nil,
		length: 0,
	}
}

func (stack *Stack[T]) Push(value T) {
	node := newStackNode[T](value)
	node.next = stack.top
	stack.top = node
	stack.length++
}

func (stack *Stack[T]) Pop() *T {
	if stack.length == 0 {
		return nil
	}

	value := &stack.top.value
	stack.top = stack.top.next
	stack.length--
	return value
}
