package ch5

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
