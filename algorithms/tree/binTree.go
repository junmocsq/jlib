package tree

import "fmt"

type arrayBinTree struct {
	arr    []interface{}
	length int
}

func NewArrayBinTree() *arrayBinTree {
	return &arrayBinTree{
		arr:    nil,
		length: 0,
	}
}
func (a *arrayBinTree) Add(ele interface{}) {
	a.arr = append(a.arr, ele)
	a.length++
}

func (a *arrayBinTree) left(index int) int {
	return index*2 + 1
}

func (a *arrayBinTree) right(index int) int {
	return index*2 + 2
}

func (a *arrayBinTree) parent(index int) int {
	return (index - 1) / 2
}

func (a *arrayBinTree) prePrint() string {
	var p func(i int)
	var ss string
	p = func(i int) {
		if i >= a.length {
			return
		}
		ss += fmt.Sprintf("%v->", a.arr[i])
		//fmt.Printf("%v->", a.arr[i])
		p(a.left(i))
		p(a.right(i))
	}
	p(0)
	//fmt.Println()
	return ss
}

func (a *arrayBinTree) midPrint() string {
	var p func(i int)
	var ss string
	p = func(i int) {
		if i >= a.length {
			return
		}
		p(a.left(i))
		ss += fmt.Sprintf("%v->", a.arr[i])
		//fmt.Printf("%v->", a.arr[i])
		p(a.right(i))
	}
	p(0)
	//fmt.Println()
	return ss
}

func (a *arrayBinTree) postPrint() string {
	var p func(i int)
	var ss string
	p = func(i int) {
		if i >= a.length {
			return
		}
		p(a.left(i))
		p(a.right(i))
		ss += fmt.Sprintf("%v->", a.arr[i])
		//fmt.Printf("%v->", a.arr[i])
	}
	p(0)

	//fmt.Println()
	return ss
}

type linkBinNode struct {
	val         interface{}
	left, right *linkBinNode
	deep        int
}

type linkBinTree struct {
	root *linkBinNode
	num  int
	deep int
}

func NewLinkBinTree() *linkBinTree {
	return &linkBinTree{
		root: nil,
		num:  0,
		deep: -1,
	}
}

func (l *linkBinTree) Add(ele interface{}) bool {
	node := &linkBinNode{
		val:   ele,
		left:  nil,
		right: nil,
		deep:  0,
	}
	if l.root == nil {
		l.root = node
		l.num++
		l.deep++
		return true
	}

	temp := l.root
	for {
		compareVal := l.compare(temp.val, ele)
		if compareVal == -999 {
			return false
		}
		if compareVal == 0 {
			return false
		}
		if compareVal == -1 { // right
			if temp.right != nil {
				temp = temp.right
			} else {
				node.deep = temp.deep + 1
				temp.right = node
				if node.deep > l.deep {
					l.deep = node.deep
				}
				return true
			}
		}
		if compareVal == 1 { // left
			if temp.left != nil {
				temp = temp.left
			} else {
				node.deep = temp.deep + 1
				temp.left = node
				if node.deep > l.deep {
					l.deep = node.deep
				}
				return true
			}
		}
	}

}

func (l *linkBinTree) Find(ele interface{}) *linkBinNode {
	temp := l.root
	if temp == nil {
		return nil
	}
	for {
		compareVal := l.compare(temp.val, ele)
		if compareVal == -999 {
			return nil
		}
		if compareVal == 0 {
			return temp
		}
		if compareVal == -1 { // right
			if temp.right != nil {
				temp = temp.right
			} else {
				return nil
			}
		}
		if compareVal == 1 { // left
			if temp.left != nil {
				temp = temp.left
			} else {
				return nil
			}
		}
	}
}

func (l *linkBinTree) prePrint() string {
	var p func(node *linkBinNode)
	var ss string
	p = func(node *linkBinNode) {
		if node == nil {
			return
		}
		ss += fmt.Sprintf("%v->", node.val)
		//fmt.Printf("%v d:%d->", node.val,node.deep)
		p(node.left)
		p(node.right)
	}
	p(l.root)
	//fmt.Println()
	return ss
}

func (l *linkBinTree) midPrint() string {
	var p func(node *linkBinNode)
	var ss string
	p = func(node *linkBinNode) {
		if node == nil {
			return
		}

		p(node.left)
		ss += fmt.Sprintf("%v->", node.val)
		//fmt.Printf("%v d:%d->", node.val,node.deep)
		p(node.right)
	}
	p(l.root)
	//fmt.Println()
	return ss
}

func (l *linkBinTree) postPrint() string {
	var p func(node *linkBinNode)
	var ss string
	p = func(node *linkBinNode) {
		if node == nil {
			return
		}
		p(node.left)
		p(node.right)
		ss += fmt.Sprintf("%v->", node.val)
		//fmt.Printf("%v d:%d->", node.val,node.deep)
	}
	p(l.root)
	//fmt.Println()
	return ss
}

// 1 大于 0 等于 -1 小于 -999 不能比较
func (l *linkBinTree) compare(e1, e2 interface{}) int {
	switch e1 := e1.(type) {
	case int:
		if e2, ok := e2.(int); ok {
			if e1 > e2 {
				return 1
			} else if e1 < e2 {
				return -1
			} else {
				return 0
			}
		} else {
			return -999
		}
	case string:
		if e2, ok := e2.(string); ok {
			if e1 > e2 {
				return 1
			} else if e1 < e2 {
				return -1
			} else {
				return 0
			}
		} else {
			return -999
		}
	default:
		return -999
	}
}
