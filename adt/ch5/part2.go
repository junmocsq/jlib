package ch5

import (
	"fmt"
)

type binNode struct {
	val   int
	left  *binNode
	right *binNode
}

type binTree struct {
	root *binNode
}

func NewBinTree() *binTree {
	return &binTree{}
}

func (b *binTree) IsEmpty() bool {
	return b.root == nil
}

func (b *binTree) MakeBT(left, right *binNode, item int) *binNode {
	binNode := &binNode{
		val:   item,
		left:  left,
		right: right,
	}
	return binNode
}

func (b *binTree) LChild(node *binNode) *binNode {
	return node.left
}

func (b *binTree) RChild(node *binNode) *binNode {
	return node.right
}

func (b *binTree) Data(node *binNode) int {
	return node.val
}

func (b *binTree) Add(val int) {
	ele := &binNode{
		val: val,
	}
	if b.root == nil {
		b.root = ele
		return
	}
	temp := b.root
	for {
		if temp.val > val {
			if temp.left == nil {
				temp.left = ele
				return
			} else {
				temp = temp.left
			}
		} else {
			if temp.right == nil {
				temp.right = ele
				return
			} else {
				temp = temp.right
			}
		}
	}
}

func (b *binTree) InOrder() {
	var f func(node *binNode)
	f = func(node *binNode) {
		if node == nil {
			return
		}
		f(node.left)
		fmt.Printf("%d ", node.val)
		f(node.right)
	}
	f(b.root)
	fmt.Println()
}

func (b *binTree) IterInOrder() {
	temp := b.root
	stack := NewBinNodeStack()
	for {
		for temp != nil {
			stack.Push(temp)
			temp = temp.left
		}
		if stack.empty() {
			break
		}
		temp = stack.Pop()
		fmt.Printf("%d ", temp.val)
		temp = temp.right
	}
	fmt.Println()
}

func (b *binTree) PreOrder() {
	var f func(node *binNode)
	f = func(node *binNode) {
		if node == nil {
			return
		}
		fmt.Printf("%d ", node.val)
		f(node.left)
		f(node.right)
	}
	f(b.root)
	fmt.Println()
}
func (b *binTree) IterPreOrder() {
	temp := b.root
	stack := NewBinNodeStack()
	for {
		for temp != nil {
			fmt.Printf("%d ", temp.val)
			stack.Push(temp)
			temp = temp.left
		}
		if stack.empty() {
			break
		}
		temp = stack.Pop()
		temp = temp.right
	}
	fmt.Println()
}
func (b *binTree) IterPreOrder2() {
	temp := b.root
	stack := NewBinNodeStack()
	stack.Push(temp)
	for !stack.empty() {
		temp = stack.Pop()
		fmt.Printf("%d ", temp.val)
		stack.Push(temp.right)
		stack.Push(temp.left)
	}
	fmt.Println()
}

func (b *binTree) PostOrder() {
	var f func(node *binNode)
	f = func(node *binNode) {
		if node == nil {
			return
		}
		f(node.left)
		f(node.right)
		fmt.Printf("%d ", node.val)
	}
	f(b.root)
	fmt.Println()
}
func (b *binTree) IterPostOrder() {
	temp := b.root
	arr := []int{}
	stack := NewBinNodeStack()
	stack.Push(temp)
	for !stack.empty() {
		temp = stack.Pop()
		stack.Push(temp.left)
		stack.Push(temp.right)
		arr = append(arr, temp.val)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func (b *binTree) LevelOrder() {
	queue := NewBinNodeQueue()
	queue.Push(b.root)
	for !queue.empty() {
		node := queue.Pop()
		queue.Push(node.left)
		queue.Push(node.right)
		fmt.Printf("%d ", node.val)
	}
	fmt.Println()
}

type binNodeStack struct {
	arr []*binNode
}

func NewBinNodeStack() *binNodeStack {
	return &binNodeStack{}
}
func (b *binNodeStack) empty() bool {
	return len(b.arr) == 0
}
func (b *binNodeStack) Push(node *binNode) {
	if node == nil {
		return
	}
	b.arr = append(b.arr, node)
}

func (b *binNodeStack) Pop() *binNode {
	if b.empty() {
		return nil
	}

	node := b.arr[len(b.arr)-1]
	b.arr = b.arr[:len(b.arr)-1]
	return node
}
func (b *binNodeStack) Top() *binNode {
	if b.empty() {
		return nil
	}
	return b.arr[len(b.arr)-1]
}

type binNodeQueue struct {
	arr []*binNode
}

func NewBinNodeQueue() *binNodeQueue {
	return &binNodeQueue{}
}
func (b *binNodeQueue) empty() bool {
	return len(b.arr) == 0
}
func (b *binNodeQueue) Push(node *binNode) {
	if node == nil {
		return
	}
	b.arr = append(b.arr, node)
}

func (b *binNodeQueue) Pop() *binNode {
	if b.empty() {
		return nil
	}
	node := b.arr[0]
	b.arr = b.arr[1:]
	return node
}
