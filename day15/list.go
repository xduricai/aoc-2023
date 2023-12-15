package day15

type Node struct {
	next  *Node
	label string
	value int
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewNode(label string, value int) *Node {
	return &Node{
		next:  nil,
		label: label,
		value: value,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (list *LinkedList) Insert(label *string, value int) {
	node := NewNode(*label, value)

	if list.length == 0 {
		list.head = node
		list.tail = node
		list.length++
		return
	}
	current := list.head

	for current != nil {
		if current.label == *label {
			current.value = value
			return
		}
		current = current.next
	}

	list.tail.next = node
	list.tail = node
	list.length++
}

func (list *LinkedList) Remove(label *string) {
	if list.length == 0 {
		return
	}

	if list.head.label == *label {
		if list.length == 1 {
			list.tail = nil
		}
		list.head = list.head.next
		list.length--
		return
	}

	previous := list.head
	current := list.head.next

	for current != nil {
		if current.label == *label {
			previous.next = current.next
			list.length--

			if current == list.tail {
				list.tail = previous
			}
			return
		}
		previous = current
		current = current.next
	}
}
