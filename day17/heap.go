package day17

type Move struct {
	x   int
	y   int
	dir int
	rep int
}

type State struct {
	heat int
	move *Move
}

type MinHeap struct {
	length int
	data   []State
}

func newMove(x, y, dir, rep int) *Move {
	return &Move{
		x:   x,
		y:   y,
		dir: dir,
		rep: rep,
	}
}

func newState(heat int, move *Move) *State {
	return &State{
		heat: heat,
		move: move,
	}
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		length: 0,
		data:   []State{},
	}
}

func (heap *MinHeap) Insert(value *State) {
	heap.data = append(heap.data, *value)
	heap.heapifyUp(heap.length)
	heap.length++
}

func (heap *MinHeap) Delete() *State {
	if heap.length == 0 {
		return nil
	}

	out := heap.data[0]
	heap.length--

	if heap.length == 0 {
		heap.data = []State{}
		return &out
	}

	heap.data[0] = heap.data[heap.length]
	heap.data = heap.data[:heap.length]
	heap.heapifyDown(0)

	return &out
}

func (heap *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	pIdx := parent(idx)
	node := heap.data[idx]
	parent := heap.data[pIdx]

	if parent.heat > node.heat {
		heap.data[idx] = parent
		heap.data[pIdx] = node
		heap.heapifyUp(pIdx)
	}
}

func (heap *MinHeap) heapifyDown(idx int) {
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if idx >= heap.length || lIdx >= heap.length {
		return
	}

	node := heap.data[idx]
	lNode := heap.data[lIdx]

	if rIdx == heap.length {
		if node.heat > lNode.heat {
			heap.data[idx] = lNode
			heap.data[lIdx] = node
		}
		return
	}

	rNode := heap.data[rIdx]

	if lNode.heat > rNode.heat && node.heat > rNode.heat {
		heap.data[idx] = rNode
		heap.data[rIdx] = node
		heap.heapifyDown(rIdx)
	} else if node.heat > lNode.heat {
		heap.data[idx] = lNode
		heap.data[lIdx] = node
		heap.heapifyDown(lIdx)
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}
