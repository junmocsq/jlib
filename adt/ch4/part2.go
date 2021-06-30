package ch4

import (
	"errors"
	"fmt"
)

var (
	ErrorListEmpty   = errors.New("链表为空")
	ErrorInvalidNode = errors.New("无效节点")
)

type ListNode struct {
	val         string
	listPointer *ListNode
}

type ListHead struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewList() *ListHead {
	return &ListHead{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// problem_4
func (l *ListHead) Length() int {
	return l.length
}

func (l *ListHead) Empty() bool {
	return l.length == 0
}

func (l *ListHead) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

func (l *ListHead) InsertAfterPos(pos *ListNode, node *ListNode) {
	if pos == nil { // 首节点
		node.listPointer = l.head.listPointer
		l.head = node
	} else {
		node.listPointer = pos.listPointer
		pos.listPointer = node
	}
	if node.listPointer == nil { // node为最后一个节点
		l.tail = node
	}
	l.length += 1
}

func (l *ListHead) Add(values ...string) {
	for _, val := range values {
		node := &ListNode{
			val:         val,
			listPointer: nil,
		}
		if l.Empty() {
			l.head = node
			l.tail = node
		} else {
			l.tail.listPointer = node
			l.tail = node
		}
		l.length++
	}

}

func (l *ListHead) DeleteFromTrail(trail *ListNode) (*ListNode, error) {
	if l.Empty() {
		return nil, ErrorListEmpty
	}
	var node *ListNode
	if trail == nil { // 删除头结点
		node = l.head
		l.head = l.head.listPointer
		if l.head == nil {
			l.tail = nil
		}
	} else {
		node = trail.listPointer
		if node == nil {
			return nil, ErrorInvalidNode
		}
		trail.listPointer = node.listPointer
		if trail.listPointer == nil {
			l.tail = trail
		}
	}
	l.length--
	return node, nil
}

func (l *ListHead) DeleteOddNode() { // problem_5
	trail := l.head
	index := 1
	for trail != nil {
		if index%2 == 0 {
			l.DeleteFromTrail(trail)
		} else {
			trail = trail.listPointer
		}
		index++
	}
	l.DeleteFromTrail(nil) // 删除首节点
}

func (l *ListHead) DeleteEvenNode() { // problem_5
	trail := l.head
	index := 0
	for trail != nil {
		if index%2 == 0 {
			l.DeleteFromTrail(trail)
		} else {
			trail = trail.listPointer
		}
		index++
	}
}

func (l *ListHead) FindPre(val string) (*ListNode, bool) {
	if l.Empty() {
		return nil, false
	}
	temp := l.head
	if temp.val == val { // 首节点
		return nil, true
	}
	for temp.listPointer != nil {
		if temp.listPointer.val == val {
			return temp, true
		}
		temp = temp.listPointer
	}
	return nil, false
}

func (l *ListHead) Find(val string) (*ListNode, bool) {
	if preNode, ok := l.FindPre(val); ok {
		if preNode == nil {
			return l.head, true
		} else {
			return preNode, true
		}
	}
	return nil, false
}

func (l *ListHead) DeleteWithNode(val string) int {
	num := 0
	for {
		if trail, ok := l.FindPre(val); ok {
			_, err := l.DeleteFromTrail(trail)
			if err != nil {
				panic(ErrorInvalidNode)
			}
			num++
		} else {
			break
		}
	}
	return num
}

func (l *ListHead) Print() {
	var result []string
	temp := l.head
	fmt.Println("length:", l.length, "head:", l.head, "tail:", l.tail)
	for temp != nil {
		result = append(result, temp.val)
		fmt.Printf("%s ", temp.val)
		temp = temp.listPointer
	}
	fmt.Println()
}

func (l *ListHead) Equal(res []string) bool {
	if l.length != len(res) {
		return false
	}
	temp := l.head
	index := 0
	for temp != nil {
		if temp.val != res[index] {
			return false
		}
		index++
		temp = temp.listPointer
	}
	return true
}

func (l *ListHead) MergeSortList(other *ListHead) { // problem_6
	if other.Empty() {
		return
	}
	if l.Empty() {
		l.head = other.head
		l.tail = other.tail
		l.length = other.length
		return
	}
	othTemp := other.head
	preTemp := l.head

	if othTemp.val < preTemp.val {
		node := othTemp
		othTemp = othTemp.listPointer
		node.listPointer = preTemp
		l.head = node
		preTemp = node
	}
	for preTemp.listPointer != nil && othTemp != nil {
		if preTemp.listPointer.val > othTemp.val {
			node := othTemp
			othTemp = othTemp.listPointer
			node.listPointer = preTemp.listPointer
			preTemp.listPointer = node
			preTemp = node
		} else {
			preTemp = preTemp.listPointer
		}
	}
	if preTemp.listPointer == nil {
		if othTemp == nil {
			l.tail = preTemp
		} else {
			preTemp.listPointer = othTemp
			l.tail = other.tail
		}
	}
	l.length += other.length
	other.Clear()
}

func (l *ListHead) MergeCross(other *ListHead) { // problem_7
	if other.Empty() {
		return
	}
	if l.Empty() {
		l.head = other.head
		l.tail = other.tail
		l.length = other.length
		return
	}
	temp := l.head
	otherTemp := other.head
	for temp != nil && otherTemp != nil {
		node := otherTemp
		otherTemp = otherTemp.listPointer
		node.listPointer = temp.listPointer
		temp.listPointer = node
		if node.listPointer == nil { // l先到达结尾
			node.listPointer = otherTemp
			break
		}
		temp = node.listPointer
	}
	l.length += other.length
	other.Clear()
}

// TODO step + 右移最多step步 -左移最多step步
func (l *ListHead) Move(val string, step int) bool { // problem_8

	return true
}
