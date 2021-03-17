package link

import (
	"fmt"
	"sync"
)

type double struct {
	head   *doubleNode
	tail   *doubleNode
	length int
	mu     *sync.RWMutex
}

var _ Linker = &double{}

func NewDouble() Linker {
	return &double{
		length: 0,
		mu:     new(sync.RWMutex),
	}
}

func (d *double) last() *doubleNode {
	return d.tail
}

func (d *double) isLast(node *doubleNode) bool {
	return d.tail == node
}

func (d *double) Find(val interface{}) int {
	temp := d.head
	index := 0
	for temp != nil {
		if Equal(temp.val, val) {
			return index
		}
		temp = temp.next
		index++
	}
	return -1
}

func (d *double) InsertByIndex(index int, val interface{}) bool {
	if index > d.length {
		return false
	}
	node := &doubleNode{
		val: val,
	}

	if index == 0 {
		node.next = d.head
		if d.head != nil {
			d.head.pre = node
		} else {
			d.tail = node
		}
		d.head = node
	} else {
		temp := d.head
		for i := 1; i < index; i++ {
			temp = temp.next
		}
		node.next = temp.next
		node.pre = temp
		if temp.next == nil {
			d.tail = node
		} else {
			temp.next.pre = node
		}
		temp.next = node
	}
	d.length++
	return false
}

func (d *double) ValueOf(index int) interface{} {
	if d.length <= index {
		return nil
	}
	temp := d.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (d *double) Add(values ...interface{}) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, val := range values {
		node := &doubleNode{
			val: val,
		}
		if d.tail == nil {
			d.head = node
		} else {
			node.pre = d.tail
			d.tail.next = node
		}
		d.tail = node
		d.length++
	}
	return true
}

func (d *double) Del(val interface{}) bool {
	if d.Empty() {
		return false
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	temp := d.head
	if Equal(temp.val, val) {
		d.head = temp.next
		if temp.next == nil {
			d.tail = nil
		}
		d.length--
		return true
	}

	temp = temp.next
	for temp != nil {
		if Equal(temp.val, val) {
			temp.pre.next = temp.next
			temp.next.pre = temp.pre
			if temp.next == nil {
				d.tail = temp.pre
			}
			d.length--
			return true
		}
		temp = temp.next
	}
	return false
}

func (d *double) DelAll(val interface{}) int {
	num := 0
	if d.Empty() {
		return num
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	temp := d.head
	for {
		if Equal(temp.val, val) {
			d.head = temp.next
			d.length--
			num++
			if d.length == 0 {
				d.tail = nil
				break
			}
			temp = temp.next
		} else {
			break
		}
	}

	temp = temp.next
	for temp != nil {
		if Equal(temp.val, val) {
			temp.pre.next = temp.next
			temp.next.pre = temp.pre
			if temp.next == nil {
				d.tail = temp.pre
			}
			d.length--
			num++
		} else {
			temp = temp.next
		}
	}
	return num
}
func (d *double) Empty() bool {
	return d.length == 0
}
func (d *double) Length() int {
	return d.length
}

func (d *double) Elements() []interface{} {
	arr := make([]interface{}, 0, d.length)
	temp := d.head
	for temp != nil {
		arr = append(arr, temp.val)
		temp = temp.next
	}
	return arr
}

func (d *double) Clear() {
	d.length = 0
	d.head = nil
	d.tail = nil
}

func (d *double) Print() {
	temp := d.head
	fmt.Printf("double length:%d eles:", d.Length())
	for temp != nil {
		fmt.Printf("%v ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}
