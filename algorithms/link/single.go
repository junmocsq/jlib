package link

import "sync"

type single struct {
	head   *singleNode
	length int
	mu     sync.RWMutex
}

var _ Linker = &single{}

func NewSingle() Linker {
	return &single{
		head:   nil,
		length: 0,
	}
}

func (s *single) last() *singleNode {
	if s.Empty() {
		return nil
	}
	temp := s.head
	for temp != nil {
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

	return false
}

func (s *single) Add(val interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	node := &singleNode{
		val:  val,
		next: nil,
	}
	s.length++
	if s.Empty() {
		s.head = node
		return true
	}
	last := s.last()
	last.next = node
	return false
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
