package link

import (
	"fmt"
	"sync"
)

type single struct {
	head   *singleNode
	length int
	mu     *sync.RWMutex
}

var _ Linker = &single{}

func NewSingle() Linker {
	return &single{
		head:   nil,
		length: 0,
		mu:     new(sync.RWMutex),
	}
}

func (s *single) last() *singleNode {
	if s.Empty() {
		return nil
	}
	temp := s.head
	for temp.next != nil {
		temp = temp.next
	}
	return temp
}

func (s *single) Find(val interface{}) int {
	temp := s.head
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

func (s *single) InsertByIndex(index int, val interface{}) bool {
	if index > s.length {
		return false
	}
	node := &singleNode{
		val: val,
	}
	if index == 0 {
		node.next = s.head
		s.head = node
	} else {
		temp := s.head
		for i := 1; i < index; i++ {
			temp = temp.next
		}
		node.next = temp.next
		temp.next = node
	}
	s.length++
	return false
}

func (s *single) ValueOf(index int) interface{} {
	if s.length <= index {
		return nil
	}
	temp := s.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.val
}

func (s *single) Add(values ...interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	last := s.last()

	for _, val := range values {
		node := &singleNode{
			val:  val,
			next: nil,
		}
		if last == nil {
			s.head = node
		} else {
			last.next = node
		}
		last = node
		s.length++
	}
	return true
}

func (s *single) Del(val interface{}) bool {
	if s.Empty() {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	temp := s.head
	if Equal(temp.val, val) {
		s.head = temp.next
		s.length--
		return true
	}

	for temp.next != nil {
		if Equal(temp.next.val, val) {
			temp.next = temp.next.next
			s.length--
			return true
		}
		temp = temp.next
	}
	return false
}

func (s *single) DelAll(val interface{}) int {
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
			s.length--
			num++
			if s.length == 0 {
				break
			}
			temp = temp.next
		} else {
			break
		}
	}

	for temp.next != nil {
		if Equal(temp.next.val, val) {
			temp.next = temp.next.next
			s.length--
			num++
		} else {
			temp = temp.next
		}
	}
	return num
}
func (s *single) Empty() bool {
	return s.length == 0
}
func (s *single) Length() int {
	return s.length
}

func (s *single) Elements() []interface{} {
	arr := make([]interface{}, 0, s.length)
	temp := s.head
	for temp != nil {
		arr = append(arr, temp.val)
		temp = temp.next
	}
	return arr
}

func (s *single) Clear() {
	s.length = 0
	s.head = nil
}

func (s *single) Print() {
	temp := s.head
	fmt.Printf("single length:%d eles:", s.Length())
	for temp != nil {
		fmt.Printf("%v ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}
