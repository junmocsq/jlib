package ch4

import "fmt"

type doubleNode struct {
	val          string
	lLink, rLink *doubleNode
}

type doubleLink struct {
	head   *doubleNode
	tail   *doubleNode
	length int
}

func NewDoubleLink() *doubleLink {
	sentinel := &doubleNode{}
	head := &doubleLink{
		head: sentinel,
		tail: sentinel,
	}

	return head
}

func (d *doubleLink) Print() {
	head := d.head
	for i := 0; i < d.length; i++ {
		fmt.Printf("%s ", head.rLink.val)
		head = head.rLink
	}
	fmt.Println()
}

func (d *doubleLink) Equal(arr []string) bool {
	if len(arr) != d.length {
		return false
	}
	head := d.head
	for i := 0; i < d.length; i++ {
		if arr[i] != head.rLink.val {
			return false
		}
		head = head.rLink
	}
	return true
}

func (d *doubleLink) AddFronts(vals ...string) {
	for i := range vals {
		d.AddNodeAfterPos(d.head, vals[i])
	}

}
func (d *doubleLink) AddRears(vals ...string) {
	for i := range vals {
		d.AddNodeAfterPos(d.tail, vals[i])
	}
}

func (d *doubleLink) AddNodeAfterPos(pre *doubleNode, val string) {
	ele := &doubleNode{
		val: val,
	}
	ele.rLink = pre.rLink
	if ele.rLink == nil {
		d.tail = ele
	} else {
		ele.rLink.lLink = ele
	}
	ele.lLink = pre
	pre.rLink = ele
	d.length++
}

func (d *doubleLink) Find(val string) *doubleNode {
	head := d.head
	for head.rLink != nil {
		if head.rLink.val == val {
			return head.rLink
		}
		head = head.rLink
	}
	return nil
}

func (d *doubleLink) Delete(val string) *doubleNode {
	ele := d.Find(val)
	if ele == nil {
		return nil
	}
	if ele.rLink == nil {
		d.tail = ele.lLink
		ele.lLink.rLink = nil
	} else {
		ele.lLink.rLink = ele.rLink
		ele.rLink.lLink = ele.lLink
	}
	d.length--
	return ele
}
func (d *doubleLink) Clear() {
	d.head.rLink = nil
	d.tail = d.head
	d.length = 0
}

type doubleLoopLink struct {
	head   *doubleNode
	length int
}

func NewDoubleLoopLink() *doubleLoopLink {
	sentinel := &doubleNode{}
	sentinel.lLink = sentinel
	sentinel.rLink = sentinel
	head := &doubleLoopLink{
		head: sentinel,
	}
	return head
}

func (d *doubleLoopLink) Print() {
	head := d.head
	for i := 0; i < d.length; i++ {
		fmt.Printf("%s ", head.rLink.val)
		head = head.rLink
	}
	fmt.Println()
}

func (d *doubleLoopLink) Equal(arr []string) bool {
	if len(arr) != d.length {
		return false
	}
	head := d.head
	for i := 0; i < d.length; i++ {
		if arr[i] != head.rLink.val {
			return false
		}
		head = head.rLink
	}
	return true
}

func (d *doubleLoopLink) AddFronts(vals ...string) {
	for i := range vals {
		d.AddNodeAfterPos(d.head, vals[i])
	}

}
func (d *doubleLoopLink) AddRears(vals ...string) {
	for i := range vals {
		d.AddNodeAfterPos(d.head.lLink, vals[i])
	}
}

func (d *doubleLoopLink) AddNodeAfterPos(pre *doubleNode, val string) {
	ele := &doubleNode{
		val: val,
	}
	ele.rLink = pre.rLink
	ele.rLink.lLink = ele
	ele.lLink = pre
	pre.rLink = ele
	d.length++
}

func (d *doubleLoopLink) Find(val string) *doubleNode {
	head := d.head
	for head.rLink != nil {
		if head.rLink.val == val {
			return head.rLink
		}
		head = head.rLink
	}
	return nil
}

func (d *doubleLoopLink) Delete(val string) *doubleNode {
	ele := d.Find(val)
	if ele == nil {
		return nil
	}
	ele.lLink.rLink = ele.rLink
	ele.rLink.lLink = ele.lLink
	d.length--
	return ele
}
func (d *doubleLoopLink) Clear() {
	d.head.rLink = d.head
	d.head.lLink = d.head
	d.length = 0
}
