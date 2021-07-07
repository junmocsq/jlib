package ch5

import (
	"fmt"
)

// 线索二叉树
// 如果ptr->left 为空，则指向ptr的中序遍历的前驱节点
// 如果ptr->right 为空，则指向ptr的中序遍历的后继节点
type binThreadNode struct {
	val                     int
	left, right             *binThreadNode
	leftThread, rightThread bool
}

type binThreadTree struct {
	root *binThreadNode
}

func NewBinTreadTree() *binThreadTree {
	sentinel := &binThreadNode{}
	sentinel.leftThread = true
	sentinel.left = sentinel
	sentinel.rightThread = false
	sentinel.right = sentinel
	return &binThreadTree{
		root: sentinel,
	}
}
func (b *binThreadTree) Add(vals ...int) {
	for _, v := range vals {
		b.add(v)
	}
}
func (b *binThreadTree) add(val int) {
	parent := b.root
	if parent.leftThread {
		b.insertLeft(parent, val)
		return
	}
	parent = parent.left
	for {
		if parent.val > val {
			if !parent.leftThread {
				parent = parent.left
			} else {
				b.insertLeft(parent, val)
				return
			}
		} else {
			if !parent.rightThread {
				parent = parent.right
			} else {
				b.insertRight(parent, val)
				return
			}
		}
	}
}

// tree 的后继节点
func (b *binThreadTree) insucc(tree *binThreadNode) *binThreadNode {
	temp := tree.right
	if !tree.rightThread {
		for !temp.leftThread {
			temp = temp.left
		}
	}
	return temp
}

// tree 的前驱节点
func (b *binThreadTree) inPre(tree *binThreadNode) *binThreadNode {
	temp := tree.right
	if !tree.leftThread {
		for !temp.rightThread {
			temp = temp.right
		}
	}
	return temp
}

func (b *binThreadTree) insertRight(parent *binThreadNode, val int) bool {
	if parent == nil {
		return false
	}
	node := &binThreadNode{
		val:         val,
		left:        parent,
		leftThread:  true,
		right:       parent.right,
		rightThread: parent.rightThread,
	}
	parent.rightThread = false
	parent.right = node
	if !node.rightThread {
		temp := b.insucc(node.right)
		temp.left = node
	}
	//fmt.Printf("right node:%v left:%v right:%v \n", node, node.left, node.right)
	return true
}

func (b *binThreadTree) insertLeft(parent *binThreadNode, val int) bool {
	if parent == nil {
		return false
	}
	node := &binThreadNode{
		val:         val,
		left:        parent.left,
		leftThread:  parent.leftThread,
		right:       parent,
		rightThread: true,
	}
	parent.leftThread = false
	parent.left = node
	if !node.leftThread { // 非叶子节点
		temp := b.inPre(node.left)
		temp.right = node
	}
	//fmt.Printf("left node:%v left:%v right:%v \n", node, node.left, node.right)
	return true
}

func (b *binThreadTree) ThreadInOrderPrint() {
	fmt.Println(b.ThreadInOrder())
}

func (b *binThreadTree) ThreadInOrder() []int {
	var arr []int
	temp := b.root.left
	if temp == nil {
		return arr
	}
	for !temp.leftThread {
		temp = temp.left
	}
	for {
		arr = append(arr, temp.val)
		temp = b.insucc(temp)
		if temp == b.root {
			break
		}
	}
	temp = b.root.left
	return arr
}
