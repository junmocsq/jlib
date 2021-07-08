package ch5

var maxCompare = func(a, b int) bool {
	return a > b
}
var minCompare = func(a, b int) bool {
	return a < b
}

type heap struct {
	arr     []int
	compare func(a1, a2 int) bool
}

func NewHeap(com func(a1, a2 int) bool) *heap {
	return &heap{
		compare: com,
	}
}
func (h *heap) Clear() {
	h.arr = nil
}
func (h *heap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *heap) leftIndex(index int) int {
	return 2*index + 1
}

func (h *heap) rightIndex(index int) int {
	return 2*index + 2
}

func (h *heap) append(item int) int {
	h.arr = append(h.arr, item)
	return len(h.arr) - 1
}

func (h *heap) deleteTail() bool {
	if h.empty() {
		return false
	}
	h.arr = h.arr[:len(h.arr)-1]
	return true
}

func (h *heap) empty() bool {
	return len(h.arr) == 0
}
func (h *heap) Add(items ...int) {
	for _, v := range items {
		h.add(v)
	}
}
func (h *heap) add(item int) {
	index := h.append(item)

	for index > 0 {
		parent := h.parentIndex(index)
		if h.compare(item, h.arr[parent]) {
			h.arr[index] = h.arr[parent]
			index = parent
		} else {
			break
		}
	}
	h.arr[index] = item
}

func (h *heap) Delete() (int, bool) {
	if h.empty() {
		return 0, false
	}
	node := h.arr[0]
	length := len(h.arr)
	h.arr[0] = h.arr[length-1]
	h.deleteTail()
	h.down(0)
	return node, true
}

// problem_3
func (h *heap) Change(index int, val int) {
	h.arr[index] = val
	h.up(index)
	h.down(index)
}

func (h heap) up(index int) {
	for index > 0 {
		parent := h.parentIndex(index)
		if h.compare(h.arr[index], h.arr[parent]) {
			h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
			index = parent
		} else {
			break
		}
	}
}

func (h heap) down(index int) {
	length := len(h.arr)
	var left, right int
	for {
		left = h.leftIndex(index)
		if left >= length {
			break
		}
		right = left + 1
		if right >= length {
			if h.compare(h.arr[left], h.arr[index]) {
				h.arr[left], h.arr[index] = h.arr[index], h.arr[left]
				index = left
			}
			break
		}
		if h.compare(h.arr[left], h.arr[right]) {
			if h.compare(h.arr[left], h.arr[index]) {
				h.arr[left], h.arr[index] = h.arr[index], h.arr[left]
				index = left
			} else {
				break
			}
		} else {
			if h.compare(h.arr[right], h.arr[index]) {
				h.arr[right], h.arr[index] = h.arr[index], h.arr[right]
				index = right
			} else {
				break
			}
		}
	}
}

// problem_4
func (h *heap) DeleteWithIndex(index int) (int, bool) {
	length := len(h.arr)
	if length <= index {
		return 0, false
	}
	node := h.arr[index]
	h.arr[index] = h.arr[length-1]
	h.deleteTail()
	h.up(index)
	h.down(index)
	return node, true
}

// problem_5
func (h *heap) Find(val int) (int, bool) {
	for k, v := range h.arr {
		if v == val {
			return k, true
		}
	}
	return 0, false
}

// -------------------------------------problem 6--------------------------------------------------------------
// 每个节点有四个指针，左右儿子和父节点指针，next为当前节点的层次遍历的下一个节点指针
type headNode struct {
	val                 int
	left, right, parent *headNode
	next                *headNode
}
type linkHeap struct {
	length  int
	root    *headNode
	compare func(a1, a2 int) bool
}

func NewLinkHeap(compare func(a1, a2 int) bool) *linkHeap {
	return &linkHeap{
		compare: compare,
	}
}

func (l *linkHeap) up(node *headNode) {
	for node.parent != nil {
		if l.compare(node.val, node.parent.val) {
			node.val, node.parent.val = node.parent.val, node.val
			node = node.parent
		} else {
			break
		}
	}
}

func (l *linkHeap) down(node *headNode) {
	for node.left != nil {
		if node.right == nil {
			if l.compare(node.left.val, node.val) {
				node.val, node.left.val = node.left.val, node.val
				node = node.left
			}
			break
		}
		if l.compare(node.left.val, node.right.val) {
			if l.compare(node.left.val, node.val) {
				node.val, node.left.val = node.left.val, node.val
				node = node.left
			} else {
				break
			}
		} else {
			if l.compare(node.right.val, node.val) {
				node.val, node.right.val = node.right.val, node.val
				node = node.right
			} else {
				break
			}
		}
	}
}

func (l *linkHeap) insertTail(val int) *headNode {
	node := &headNode{
		val: val,
	}
	if l.length == 0 {
		l.root = node
		l.length++
		return node
	}
	tailParent := l.root
	for i := 0; i < (l.length-1)/2; i++ {
		tailParent = tailParent.next
	}
	if tailParent.left == nil {
		tailParent.left = node
	} else {
		tailParent.right = node
	}
	node.parent = tailParent
	for tailParent.next != nil {
		tailParent = tailParent.next
	}
	tailParent.next = node
	l.length++
	return node
}

func (l *linkHeap) deleteTail() *headNode {
	if l.length == 0 {
		return nil
	}
	var node *headNode
	if l.length == 1 {
		node = l.root
		l.root = nil
		l.length--
		return node
	}
	tailParent := l.root
	for i := 0; i < (l.length-2)/2; i++ {
		tailParent = tailParent.next
	}
	if tailParent.right != nil {
		node = tailParent.right
		tailParent.right = nil
	} else {
		node = tailParent.left
		tailParent.left = nil
	}
	for tailParent.next != node {
		tailParent = tailParent.next
	}
	tailParent.next = nil
	l.length--
	return node
}

func (l *linkHeap) Add(items ...int) {
	for _, v := range items {
		l.add(v)
	}
}
func (l *linkHeap) add(item int) {
	node := l.insertTail(item)
	l.up(node)
}

func (l *linkHeap) Delete() (int, bool) {
	if l.length == 0 {
		return 0, false
	}
	node := l.deleteTail()
	if l.length == 0 {
		return node.val, true
	}
	result := l.root.val
	l.root.val = node.val
	l.down(l.root)
	return result, true
}

func (l *linkHeap) Arr() []int {
	var arr []int
	temp := l.root
	for temp != nil {
		arr = append(arr, temp.val)
		temp = temp.next
	}
	return arr
}
