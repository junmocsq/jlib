package link

import (
	"fmt"
	"sync"
)

type doubleCircular struct {
	head   *doubleNode
	length int
	mu     *sync.RWMutex
}

var _ Linker = &doubleCircular{}

func NewDoubleCircular() Linker {
	return &doubleCircular{
		head:   nil,
		length: 0,
		mu:     new(sync.RWMutex),
	}
}

func (d *doubleCircular) last() *doubleNode {
	if d.Empty() {
		return nil
	}
	return d.head.pre
}

func (d *doubleCircular) checkIsLast(node *doubleNode) bool {
	return node == d.last()
}

func (d *doubleCircular) Find(val interface{}) int {
	temp := d.head
	if temp == nil {
		return -1
	}

	for index := 0; index < d.length; index++ {
		if Equal(temp.val, val) {
			return index
		}
		temp = temp.next
	}
	return -1
}

func (d *doubleCircular) FindAll(val interface{}) []int {
	var res []int
	temp := d.head

	for index := 0; index < d.length; index++ {
		if Equal(temp.val, val) {
			res = append(res, index)
		}
		temp = temp.next
	}
	return res
}

func (d *doubleCircular) InsertByIndex(index int, val interface{}) bool {

	if index > d.length {
		return false
	}
	node := &doubleNode{
		val: val,
	}
	if d.Empty() {
		node.next = node
		node.pre = node
		d.head = node
		d.length++
		return true
	}

	temp := d.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	node.next = temp
	node.pre = temp.pre
	temp.pre = node
	node.pre.next = node
	if index == 0 {
		d.head = node
	}
	d.length++
	return true
}

func (d *doubleCircular) ValueOf(index int) interface{} {
	if d.length <= index {
		return nil
	}
	temp := d.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (d *doubleCircular) Add(values ...interface{}) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	var tail *doubleNode
	if d.head == nil {
		tail = nil
	} else {
		tail = d.head.pre
	}

	for _, val := range values {
		node := &doubleNode{
			val: val,
		}
		if tail == nil {
			node.next = node
			node.pre = node
			d.head = node
		} else {
			node.next = tail.next
			node.pre = tail
			tail.next = node
			node.next.pre = node
		}
		tail = node
		d.length++
	}

	return true
}

func (d *doubleCircular) Del(val interface{}) bool {
	if d.Empty() {
		return false
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	temp := d.head
	length := d.length
	for i := 0; i < length; i++ {
		if Equal(temp.val, val) {
			temp.next.pre = temp.pre
			temp.pre.next = temp.next
			d.length--

			if temp == d.head {
				d.head = temp.next
			}

			if d.length == 0 {
				d.head = nil
			}
			return true
		}

		temp = temp.next
	}
	return false
}

func (d *doubleCircular) DelHead() interface{} {
	if d.length == 0 {
		return nil
	}
	temp := d.head

	temp.next.pre = temp.pre
	temp.pre.next = temp.next
	d.length--
	d.head = temp.next

	if d.length == 0 {
		d.head = nil
	}

	return temp.val
}

func (d *doubleCircular) DelTail() interface{} {
	if d.length == 0 {
		return nil
	}
	temp := d.head.pre

	temp.next.pre = temp.pre
	temp.pre.next = temp.next
	d.length--
	if d.length == 0 {
		d.head = nil
	}
	return temp.val
}

func (d *doubleCircular) DelByIndex(index int) interface{} {
	if d.length <= index || index < 0 {
		return nil
	}
	var val interface{}
	temp := d.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	val = temp.val
	temp.next.pre = temp.pre
	temp.pre.next = temp.next
	d.length--

	if temp == d.head {
		d.head = temp.next
	}

	if d.length == 0 {
		d.head = nil
	}
	return val
}

func (d *doubleCircular) DelAll(val interface{}) int {
	num := 0
	if d.Empty() {
		return num
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	temp := d.head
	length := d.length
	for i := 0; i < length; i++ {
		if Equal(temp.val, val) {
			temp.next.pre = temp.pre
			temp.pre.next = temp.next
			d.length--
			num++
			if temp == d.head {
				d.head = temp.next
			}

			if d.length == 0 {
				d.head = nil
				break
			}
		}
		temp = temp.next
	}

	return num
}
func (d *doubleCircular) Empty() bool {
	return d.length == 0
}
func (d *doubleCircular) Length() int {
	return d.length
}

func (d *doubleCircular) Elements() []interface{} {
	arr := make([]interface{}, 0, d.length)
	temp := d.head
	for temp != nil {
		arr = append(arr, temp.val)
		if d.checkIsLast(temp) {
			break
		}
		temp = temp.next
	}
	return arr
}

func (d *doubleCircular) Clear() {
	d.length = 0
	d.head = nil
}

func (d *doubleCircular) Print() {
	temp := d.head
	fmt.Printf("double circular length:%d eles:", d.Length())

	for index := 0; index < d.length; index++ {
		//time.Sleep(time.Second)
		fmt.Printf("%v ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}
