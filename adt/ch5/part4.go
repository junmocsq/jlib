package ch5

// 复制 5.4
func (b *binTree) Copy() *binTree {
	newTree := NewBinTree()
	var tcopy func(node *binNode) *binNode
	tcopy = func(node *binNode) *binNode {
		if node == nil {
			return nil
		}
		temp := &binNode{
			val: node.val,
		}
		temp.left = tcopy(node.left)
		temp.right = tcopy(node.right)
		return temp
	}
	newTree.root = tcopy(b.root)
	return newTree
}

// 判等
func (b *binTree) Equal(other *binTree) bool {
	var equal func(first, second *binNode) bool
	equal = func(first, second *binNode) bool {
		if first == nil || second == nil {
			if first == nil && second == nil {
				return true
			} else {
				return false
			}
		}
		if first.val != second.val {
			return false
		}
		if !equal(first.left, second.left) {
			return false
		}
		if !equal(first.right, second.right) {
			return false
		}
		return true
	}
	return equal(b.root, other.root)
}

// problem_1
func (b *binTree) LeafNum() int {
	var num int
	var leafNum func(node *binNode)
	leafNum = func(node *binNode) {
		if node == nil {
			return
		}
		if node.left == nil && node.right == nil {
			num++
			return
		}
		leafNum(node.left)
		leafNum(node.right)
	}
	leafNum(b.root)
	return num
}

func (b *binTree) SwapNode() {
	var swapNode func(parent *binNode)
	swapNode = func(parent *binNode) {
		if parent == nil {
			return
		}
		parent.left, parent.right = parent.right, parent.left
		swapNode(parent.left)
		swapNode(parent.right)
	}
	swapNode(b.root)
}
