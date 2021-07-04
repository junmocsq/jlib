package ch4

import "fmt"

// type ListNode struct {
//	val         string
//	listPointer *ListNode
//}
// 翻转单链表
func (l *ListHead) Reverse() {
	lead := l.head
	var trail, middle *ListNode
	for lead != nil {
		trail = middle
		middle = lead
		lead = lead.listPointer
		middle.listPointer = trail
	}
	l.head, l.tail = l.tail, l.head
}

// 两个链表相连
func (l *ListHead) Concatenate(l2 *ListHead) {
	if l.Empty() {
		l.head = l2.head
		l.tail = l2.tail
	} else {
		l.tail.listPointer = l2.head
	}
	l.length += l2.length
	l2.Clear()
}

type loopNode struct {
	val  string
	link *loopNode
}

type loopLink struct {
	tail   *loopNode
	length int
}

func NewLoop() *loopLink {
	sentinel := &loopNode{} // 哨兵节点
	sentinel.link = sentinel
	return &loopLink{
		tail:   sentinel,
		length: 0,
	}
}

func (l *loopLink) Empty() bool {
	return l.length == 0
}

func (l *loopLink) AddRear(val string) {
	ele := &loopNode{
		val:  val,
		link: l.tail.link,
	}
	l.tail.link = ele
	l.tail = ele
	l.length++
}
func (l *loopLink) AddRears(vals ...string) {
	for i := range vals {
		l.AddRear(vals[i])
	}
}
func (l *loopLink) AddFront(val string) {
	head := l.tail.link
	ele := &loopNode{
		val:  val,
		link: head.link,
	}
	head.link = ele
	if l.Empty() {
		l.tail = ele
	}
	l.length++
}
func (l *loopLink) AddFronts(vals ...string) {
	for i := range vals {
		l.AddFront(vals[i])
	}
}

func (l *loopLink) Print() {
	pre := l.tail.link
	for i := 0; i < l.length; i++ {
		fmt.Printf("-%s ", pre.link.val)
		pre = pre.link
	}
	fmt.Println()
}
func (l *loopLink) Clear() {
	l.length = 0
	l.tail = l.tail.link
	l.tail.link = l.tail
}
func (l *loopLink) Equal(vals []string) bool {
	if l.length != len(vals) {
		return false
	}
	pre := l.tail.link
	for i := 0; i < l.length; i++ {
		if pre.link.val != vals[i] {
			return false
		}
		pre = pre.link
	}
	return true
}

func (l *loopLink) Find(val string) *loopNode { // problem_1
	pre := l.findPre(val)
	if pre == nil {
		return nil
	}
	return pre.link
}

func (l *loopLink) findPre(val string) *loopNode { // problem_1
	head := l.tail.link
	for i := 0; i < l.length; i++ {
		if head.link.val == val {
			return head
		}
	}
	return nil
}

func (l *loopLink) Delete(val string) *loopNode { // problem_2
	pre := l.findPre(val)
	var node *loopNode
	if pre != nil {
		node = pre.link
		pre.link = pre.link.link
		l.length--
	}
	return node
}

func (l *loopLink) Concatenate(l2 *loopLink) { // problem_3
	if l2.Empty() {
		return
	}
	head1 := l.tail.link
	head2 := l2.tail.link
	l.tail.link = l2.tail.link.link
	l.tail = l2.tail
	l.tail.link = head1
	l.length += l2.length

	// l2 清空
	l2.length = 0
	l2.tail = head2
	head2.link = l2.tail
}

// 反转 画图
func (l *loopLink) Reverse() { // problem_3
	t := l.tail.link.link
	lead := l.tail.link
	var trail, middle *loopNode
	middle = l.tail
	for i := 0; i < l.length+1; i++ {
		trail = middle
		middle = lead
		lead = lead.link
		middle.link = trail
	}
	l.tail = t
}
