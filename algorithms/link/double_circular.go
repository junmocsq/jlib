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

func (s *doubleCircular) last() *doubleNode {
	if s.Empty() {
		return nil
	}
	return s.head.pre
}

func (s *doubleCircular) checkIsLast(node *doubleNode) bool {
	return node == s.last()
}

func (s *doubleCircular) Find(val interface{}) int {
	temp := s.head
	if temp == nil {
		return -1
	}

	for index := 0; index < s.length; index++ {
		if Equal(temp.val, val) {
			return index
		}
		temp = temp.next
	}
	return -1
}

func (s *doubleCircular) InsertByIndex(index int, val interface{}) bool {

	if index > s.length {
		return false
	}
	node := &doubleNode{
		val: val,
	}
	if s.Empty() {
		node.next = node
		node.pre = node
		s.head = node
		s.length++
		return true
	}

	temp := s.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	node.next = temp
	node.pre = temp.pre
	temp.pre = node
	node.pre.next = temp
	if index == 0 {
		s.head = node
	}
	s.length++
	return true
}

func (s *doubleCircular) ValueOf(index int) interface{} {
	if s.length <= index {
		return nil
	}
	temp := s.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (s *doubleCircular) Add(values ...interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	var tail *doubleNode
	if s.head == nil {
		tail = nil
	} else {
		tail = s.head.pre
	}

	for _, val := range values {
		node := &doubleNode{
			val: val,
		}
		if tail == nil {
			node.next = node
			node.pre = node
			s.head = node
		} else {
			node.next = tail.next
			node.pre = tail
			tail.next = node
			node.next.pre = node

			tail = node
		}

		s.length++
	}
	return true
}

func (s *doubleCircular) Del(val interface{}) bool {
	if s.Empty() {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	temp := s.head
	for {
		if Equal(temp.val, val) {
			s.head = temp.next
			temp.next.pre = temp.pre
			temp.pre.next = temp.next
			s.length--

			if temp == s.head {
				s.head = temp.next
			}

			if s.length == 0 {
				s.head = nil
			}

			return true
		}
		if s.checkIsLast(temp) {
			break
		}
		temp = temp.next
	}
	return false
}

func (s *doubleCircular) DelAll(val interface{}) int {
	num := 0
	if s.Empty() {
		return num
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	temp := s.head
	for {
		if Equal(temp.val, val) {
			s.head = temp.next
			temp.next.pre = temp.pre
			temp.pre.next = temp.next
			s.length--
			num++
			if temp == s.head {
				s.head = temp.next
			}

			if s.length == 0 {
				s.head = nil
				break
			}

		}
		if s.checkIsLast(temp) {
			break
		}
		temp = temp.next
	}

	return num
}
func (s *doubleCircular) Empty() bool {
	return s.length == 0
}
func (s *doubleCircular) Length() int {
	return s.length
}

func (s *doubleCircular) Elements() []interface{} {
	arr := make([]interface{}, 0, s.length)
	temp := s.head
	for temp != nil {
		arr = append(arr, temp.val)
		if s.checkIsLast(temp) {
			break
		}
		temp = temp.next
	}
	return arr
}

func (s *doubleCircular) Clear() {
	s.length = 0
	s.head = nil
}

func (s *doubleCircular) Print() {
	temp := s.head
	fmt.Printf("double circular length:%d eles:", s.Length())

	for index := 0; index < s.length; index++ {
		//time.Sleep(time.Second)
		fmt.Printf("%v ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}
