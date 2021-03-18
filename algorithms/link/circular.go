package link

import (
	"fmt"
	"sync"
)

type circular struct {
	tail   *singleNode
	length int
	mu     *sync.RWMutex
}

var _ Linker = &circular{}

func NewCircular() Linker {
	return &circular{
		tail:   nil,
		length: 0,
		mu:     new(sync.RWMutex),
	}
}

func (s *circular) last() *singleNode {
	return s.tail
}

func (s *circular) head() *singleNode {
	if s.tail == nil {
		return nil
	}
	return s.tail.next
}

func (s *circular) checkIsLast(node *singleNode) bool {
	return node == s.tail
}

func (s *circular) Find(val interface{}) int {
	temp := s.head()
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

func (s *circular) FindAll(val interface{}) []int {
	var res []int
	temp := s.head()
	if temp == nil {
		return res
	}

	for index := 0; index < s.length; index++ {
		if Equal(temp.val, val) {
			res = append(res, index)
		}
		temp = temp.next
	}
	return res
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
		s.length++
		return true
	}

	pre := s.tail
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	node.next = pre.next
	pre.next = node
	if s.length == index { // 新节点添加到尾巴
		s.tail = node
	}
	s.length++
	return true
}

func (s *circular) ValueOf(index int) interface{} {
	if s.length <= index {
		return nil
	}
	temp := s.head()
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
		} else {
			node.next = s.tail.next
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

	pre := s.tail
	for i := 0; i < s.length; i++ {
		if Equal(pre.next.val, val) {
			if s.checkIsLast(pre.next) {
				s.tail = pre
			}
			pre.next = pre.next.next
			s.length--
			if s.length == 0 {
				s.tail = nil
			}
			return true
		}
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

	pre := s.tail
	length := s.length
	for i := 0; i < length; i++ {
		if Equal(pre.next.val, val) {
			if s.checkIsLast(pre.next) {
				s.tail = pre
			}
			pre.next = pre.next.next
			s.length--
			num++
			if s.length == 0 {
				s.tail = nil
			}
		} else {
			pre = pre.next
		}
	}
	return num
}

func (s *circular) DelHead() interface{} {

	return s.DelByIndex(0)
}

func (s *circular) DelTail() interface{} {

	return s.DelByIndex(s.length - 1)
}

func (s *circular) DelByIndex(index int) interface{} {
	if s.length <= index || index < 0 {
		return nil
	}
	var val interface{}
	pre := s.tail
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	val = pre.next.val
	pre.next = pre.next.next
	s.length--
	if s.checkIsLast(pre.next) {
		s.tail = pre
	}
	if s.length == 0 {
		s.tail = nil
	}
	return val
}

func (s *circular) Empty() bool {
	return s.length == 0
}
func (s *circular) Length() int {
	return s.length
}

func (s *circular) Elements() []interface{} {
	arr := make([]interface{}, 0, s.length)
	temp := s.head()
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
	s.tail = nil
}

func (s *circular) Print() {
	temp := s.head()
	fmt.Printf("circular length:%d eles:", s.Length())

	for index := 0; index < s.length; index++ {
		//time.Sleep(time.Second)
		fmt.Printf("%v ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}
