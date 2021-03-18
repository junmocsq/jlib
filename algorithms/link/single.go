package link

import (
	"fmt"
	"sync"
)

type single struct {
	head   *singleNode
	tail   *singleNode
	length int
	mu     *sync.RWMutex
}

var _ Linker = &single{}

func NewSingle() Linker {
	return &single{
		length: 0,
		mu:     new(sync.RWMutex),
	}
}

func (s *single) last() *singleNode {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.tail
}

func (s *single) checkIsLast(node *singleNode) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.tail == node
}

func (s *single) Find(val interface{}) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
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

func (s *single) FindAll(val interface{}) []int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var res []int
	temp := s.head
	index := 0
	for temp != nil {
		if Equal(temp.val, val) {
			res = append(res, index)
		}
		temp = temp.next
		index++
	}
	return res
}

func (s *single) InsertByIndex(index int, val interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index > s.length {
		return false
	}
	node := &singleNode{
		val: val,
	}
	if index == 0 {
		node.next = s.head
		if s.head == nil {
			s.tail = node
		}
		s.head = node
	} else {
		temp := s.head
		for i := 1; i < index; i++ {
			temp = temp.next
		}
		node.next = temp.next
		temp.next = node
		if s.length == index {
			s.tail = node
		}
	}
	s.length++
	return true
}

func (s *single) ValueOf(index int) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
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

	for _, val := range values {
		node := &singleNode{
			val: val,
		}
		if s.tail == nil {
			s.head = node
		} else {
			s.tail.next = node
		}
		s.tail = node
		s.length++
	}
	return true
}

func (s *single) Del(val interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.length == 0 {
		return false
	}

	temp := s.head
	if Equal(temp.val, val) {
		s.head = temp.next
		s.length--
		if s.length == 0 {
			s.tail = nil
		}
		return true
	}

	for temp.next != nil {
		if Equal(temp.next.val, val) {
			if s.checkIsLast(temp.next) {
				temp.next = nil
				s.tail = temp
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

func (s *single) DelAll(val interface{}) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	num := 0
	if s.length == 0 {
		return num
	}

	temp := s.head
	for {
		if Equal(temp.val, val) {
			s.head = temp.next
			s.length--
			num++
			if s.length == 0 {
				s.tail = nil
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
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.length == 0
}
func (s *single) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.length
}

func (s *single) Elements() []interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	arr := make([]interface{}, 0, s.length)
	temp := s.head
	for temp != nil {
		arr = append(arr, temp.val)
		temp = temp.next
	}
	return arr
}

func (s *single) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.length = 0
	s.head = nil
	s.tail = nil
}

func (s *single) Print() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	temp := s.head
	fmt.Printf("single length:%d eles:", s.Length())
	for temp != nil {
		fmt.Printf("%v ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}

func (s *single) Reverse() {
	if s.length <= 1 {
		return
	}
	pre := s.head
	next := pre.next
	pre.next = nil
	for next != nil {
		temp := next.next
		next.next = pre
		pre = next
		next = temp
	}
	s.head = pre
	return
}

// 环检测，一个走一步，一个一次走两步
func (s *single) CheckLoop() bool {
	if s.length == 0 {
		return false
	}
	t1 := s.head.next
	if t1.next == nil {
		return false
	}
	t2 := t1.next
	for t1 != nil && t2 != nil {
		if t1 == t2 {
			return true
		}
		t1 = t1.next
		if t2.next == nil {
			break
		}
		t2 = t2.next.next
	}
	return false
}

func (s *single) createLoop() int {
	num := 0
	s.Clear()
	s.Add("a")
	s.head.next = s.head
	if s.CheckLoop() {
		num++
	}
	s.Clear()
	s.Add("a", "b", "c", "d")
	s.last().next = s.head.next
	if s.CheckLoop() {
		num++
	}
	return num
}

func (s *single) DelByIndex(index int) interface{} {
	if s.length <= index || index < 0 {
		return nil
	}
	var val interface{}
	temp := s.head
	if index == 0 {
		val = temp.val
		s.head = temp.next
	} else {
		for i := 1; i < index; i++ {
			temp = temp.next
		}
		val = temp.next.val
		temp.next = temp.next.next
	}
	s.length--
	return val
}

func (s *single) DelHead() interface{} {

	return s.DelByIndex(0)
}

func (s *single) DelTail() interface{} {

	return s.DelByIndex(s.length - 1)
}

// 删除链表倒数第 n 个结点
// 		1. 循环遍历一遍，再次循环遍历找到节点删除
// 		2. 快慢指针，快指针先走n个节点，然手 快、慢指针一起走，快指针走到结尾，慢指针所在节点即为要删除的节点

func (s *single) DelDescN1(n int) *singleNode {
	length := 0
	temp := s.head
	for temp != nil {
		length++
		temp = temp.next
	}
	temp = s.head
	delIndex := length - n - 1
	if delIndex < 0 {
		return nil
	}

	var ret *singleNode
	if delIndex == 0 {
		ret = temp
		s.head = temp.next
	} else {
		for i := 0; i < delIndex-1; i++ {
			temp = temp.next
		}
		ret = temp.next
		temp.next = temp.next.next
	}
	s.length--
	return ret
}
func (s *single) DelDescN2(n int) *singleNode {
	t1 := s.head
	t2 := s.head
	var ret *singleNode
	for i := 0; i <= n; i++ {
		t2 = t2.next
		if t2 == nil {
			if i == n {
				//
				ret = s.head
				s.head = s.head.next
				s.length--
				return ret
			} else {
				return nil
			}
		}
	}

	for t2.next != nil {
		t1 = t1.next
		t2 = t2.next
	}
	ret = t1.next
	t1.next = t1.next.next
	s.length--
	return ret
}

// 求链表的中间结点:一个一次走一步，一个走两步
func (s *single) Middle() *singleNode {
	if s.length == 0 {
		return nil
	}
	t1 := s.head
	if t1.next == nil {
		return t1
	}
	t2 := t1.next.next
	if t2 == nil {
		return t1
	}
	for {
		t2 = t2.next
		if t2 == nil {
			return t1.next
		}
		t2 = t2.next
		if t2 == nil {
			return t1.next
		}
		t1 = t1.next
	}
}
