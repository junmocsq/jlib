package ch5

/**
type binNode struct {
	val   int
	left  *binNode
	right *binNode
}

type binTree struct {
	root *binNode
}
*/
func (b *binTree) Find(val int) *binNode {
	var search func(node *binNode) *binNode
	search = func(node *binNode) *binNode {
		if node == nil {
			return nil
		}
		if node.val == val {
			return node
		} else if node.val > val {
			return search(node.left)
		} else {
			return search(node.right)
		}
	}
	return search(b.root)
}

func (b *binTree) IterFind(val int) *binNode {
	node := b.root
	for node != nil {
		if node.val == val {
			return node
		} else if node.val > val {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

// pre 代表最后查找值的父级或者最后一个查找节点 bool代表是否存在查找值
func (b *binTree) findPre(val int) (*binNode, bool) {
	if b.IsEmpty() {
		return nil, false
	}
	temp := b.root
	if temp.val == val {
		return nil, true
	}
	for {
		if temp.val > val {
			if temp.left == nil {
				return temp, false
			} else if temp.left.val == val {
				return temp, true
			} else {
				temp = temp.left
			}
		} else {
			if temp.right == nil {
				return temp, false
			} else if temp.right.val == val {
				return temp, true
			} else {
				temp = temp.right
			}
		}
	}
}

func (b *binTree) Insert(val int) {
	b.InsertIter(val)
}

func (b *binTree) InsertIter(val int) {
	if b.IsEmpty() {
		b.root = &binNode{
			val:   val,
			count: 1,
		}
		return
	}
	pre := b.root
	for {
		if pre.val == val {
			pre.count++
			break
		} else if pre.val > val {
			// 左
			if pre.left == nil {
				pre.left = &binNode{
					val:   val,
					count: 1,
				}
				break
			} else {
				pre = pre.left
			}
		} else {
			// 右
			if pre.right == nil {
				pre.right = &binNode{
					val:   val,
					count: 1,
				}
				break
			} else {
				pre = pre.right
			}
		}
	}
}
func (b *binTree) Clear() {
	b.root = nil
}
func (b *binTree) InsertRecv(val int) {
	if b.IsEmpty() {
		b.root = &binNode{
			val:   val,
			count: 1,
		}
		return
	}
	pre := b.root
	var insert func(pre *binNode)
	insert = func(pre *binNode) {
		if pre.val == val {
			pre.count++
		} else if pre.val > val {
			// 左
			if pre.left == nil {
				pre.left = &binNode{
					val:   val,
					count: 1,
				}
			} else {
				insert(pre.left)
			}
		} else {
			// 右
			if pre.right == nil {
				pre.right = &binNode{
					val:   val,
					count: 1,
				}
			} else {
				insert(pre.right)
			}
		}
	}
	insert(pre)
}

// 节点删除
// 分为3种情况
// 		(1).删除节点没有子节点
//		(2).删除节点存在一个子节点
//		(3).删除节点存在两个子节点，找到节点左端的最大值节点node1，替换，如果node1存在两个节点，循环(3)，只有一个子节点，跳转(2)，没有子节点，跳转(1)
func (b *binTree) Delete(val int) bool {
	return b.DeleteIter(val)
}

func (b *binTree) DeleteIter(val int) bool {
	if b.IsEmpty() {
		return false
	}
	var pre *binNode
	node := b.root
	for {
		if node.val == val {
			if node.count > 1 {
				node.count--
				return true
			}
			if node.left != nil && node.right != nil {
				// node最大左节点
				pre = node
				swapNode := node.left
				for swapNode.right != nil {
					pre = swapNode
					swapNode = swapNode.right
				}

				node.val, swapNode.val = swapNode.val, node.val
				node.count, swapNode.count = swapNode.count, node.count
				node = swapNode
			} else if node.left != nil {
				node.val, node.left.val = node.left.val, node.val
				node.count, node.left.count = node.left.count, node.count
				node.left = nil
				return true
			} else if node.right != nil {
				node.val, node.right.val = node.right.val, node.val
				node.count, node.right.count = node.right.count, node.count
				node.right = nil
				return true
			} else {
				if pre == nil { // 根结点
					b.root = nil
				} else {
					if pre.left == node {
						pre.left = nil
					} else {
						pre.right = nil
					}
				}
				return true
			}
		} else if node.val > val {
			if node.left == nil {
				return false
			}
			pre = node
			node = node.left
		} else {
			if node.right == nil {
				return false
			}
			pre = node
			node = node.right
		}
	}
}

// 递归删除
func (b *binTree) DeleteRecv(val int) bool {
	if b.IsEmpty() {
		return false
	}
	var delete func(pre, node *binNode) bool
	delete = func(pre, node *binNode) bool {
		if node.val == val {
			if node.count > 1 {
				node.count--
				return true
			}
			if node.left != nil && node.right != nil {
				// node最大左节点
				pre = node
				swapNode := node.left
				for swapNode.right != nil {
					pre = swapNode
					swapNode = swapNode.right
				}

				node.val, swapNode.val = swapNode.val, node.val
				node.count, swapNode.count = swapNode.count, node.count
				return delete(pre, swapNode)
			} else if node.left != nil {
				node.val, node.left.val = node.left.val, node.val
				node.count, node.left.count = node.left.count, node.count
				node.left = nil
				return true
			} else if node.right != nil {
				node.val, node.right.val = node.right.val, node.val
				node.count, node.right.count = node.right.count, node.count
				node.right = nil
				return true
			} else {
				if pre == nil { // 根结点
					b.root = nil
				} else {
					if pre.left == node {
						pre.left = nil
					} else {
						pre.right = nil
					}
				}
				return true
			}
		} else if node.val > val {
			if node.left == nil {
				return false
			}
			return delete(node, node.left)
		} else {
			if node.right == nil {
				return false
			}
			return delete(node, node.right)
		}
	}
	return delete(nil, b.root)
}

func (b *binTree) findLeftMax(pre *binNode) (*binNode, *binNode) {
	node := pre.left
	for node.right != nil {
		pre = node
		node = node.right
	}
	return pre, node
}
