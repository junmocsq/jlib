package link

import (
	"fmt"
	"sync"
)

type circular struct {
	head   *singleNode
	tail   *singleNode
	length int
	mu     *sync.RWMutex
}

var _ Linker = &circular{}

func NewCircular() Linker {
	return &circular{
		head:   nil,
		tail:   nil,
		length: 0,
		mu:     new(sync.RWMutex),
	}
}

func (s *circular) last() *singleNode {
	return s.tail
}

func (s *circular) checkIsLast(node *singleNode) bool {
	return node == s.tail
}

func (s *circular) Find(val interface{}) int {
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

func (s *circular) InsertByIndex(index int, val interface{}) bool {

	if index > s.length {
		return false
	}
	node := &singleNode{
		val: val,
	}
	if s.Empty() {
		node.next = node
		s.tail = node
		s.head = node
		s.length++
		return true
	}

	if index == 0 { // 插头
		node.next = s.head
		s.head = node
		s.tail = s.head
	} else if index == s.length { // 插尾
		node.next = s.head
		s.tail.next = node
		s.tail = node
	} else { // 插中
		temp := s.head
		for i := 1; i < index; i++ {
			temp = temp.next
		}
		node.next = temp.next
		temp.next = node
	}
	s.length++
	return true
}

func (s *circular) ValueOf(index int) interface{} {
	if s.length <= index {
		return nil
	}
	temp := s.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (s *circular) Add(values ...interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, val := range values {
		node := &singleNode{
			val:  val,
			next: nil,
		}
		if s.tail == nil {
			node.next = node
			s.head = node
		} else {
			node.next = s.head
			s.tail.next = node
		}
		s.tail = node
		s.length++
	}
	return true
}

func (s *circular) Del(val interface{}) bool {
	if s.Empty() {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	defer func() {
		if s.length == 0 {
			s.head = nil
			s.tail = nil
		}
	}()
	temp := s.head
	if Equal(temp.val, val) { // 删除首元素
		s.head = temp.next
		s.tail.next = s.head
		s.length--
		return true
	}

	for index := 1; index < s.length; index++ {
		if Equal(temp.next.val, val) {
			if s.checkIsLast(temp.next) {
				s.tail = temp
				temp.next = s.head
			} else {
				temp.next = temp.next.next
			}
			s.length--
			return true
		}
		temp = temp.next
	}
	return false
}

func (s *circular) DelAll(val interface{}) int {
	num := 0
	if s.Empty() {
		return num
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	defer func() {
		if s.length == 0 {
			s.head = nil
			s.tail = nil
		}
	}()
	temp := s.head
	for { // 删除首部
		if Equal(temp.val, val) {
			s.head = temp.next
			s.tail.next = s.head
			s.length--
			num++
			temp = temp.next
		} else {
			break
		}
	}
	length := s.length
	for index := 1; index < length; index++ {
		if Equal(temp.next.val, val) {
			if s.checkIsLast(temp.next) {
				s.tail = temp
				temp.next = s.head
			} else {
				temp.next = temp.next.next
			}
			s.length--
			num++
		} else {
			temp = temp.next
		}
	}
	return num
}
func (s *circular) Empty() bool {
	return s.length == 0
}
func (s *circular) Length() int {
	return s.length
}

func (s *circular) Elements() []interface{} {
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

func (s *circular) Clear() {
	s.length = 0
	s.head = nil
	s.tail = nil
}

func (s *circular) Print() {
	temp := s.head
	fmt.Printf("circular length:%d eles:", s.Length())

	for index := 0; index < s.length; index++ {
		//time.Sleep(time.Second)
		fmt.Printf("%v ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}
