package ch4

import (
	"errors"
	"fmt"
)

var (
	ErrorEmpty = errors.New("空")
)

// 链式队列 链式栈

type Element struct {
	val  string
	next *Element
}

type stackLink struct {
	top    *Element
	length int
}

func NewStackLink() *stackLink {
	return &stackLink{}
}
func (s *stackLink) Push(val string) {
	ele := &Element{
		val:  val,
		next: nil,
	}
	ele.next = s.top
	s.top = ele
	s.length++
}

func (s *stackLink) Pop() (string, error) {
	if s.empty() {
		return "", ErrorEmpty
	}
	ele := s.top
	s.top = s.top.next
	s.length--
	return ele.val, nil
}

func (s *stackLink) empty() bool {
	return s.length == 0
}
func (s *stackLink) Length() int {
	return s.length
}

type queueLink struct {
	head   *Element
	tail   *Element
	length int
}

func NewQueueLink() *queueLink {
	return &queueLink{}
}

func (q *queueLink) lPush(val string) {
	ele := &Element{
		val:  val,
		next: nil,
	}
	if q.empty() {
		q.tail = ele
	} else {
		ele.next = q.head
	}
	q.head = ele
	q.length++
}

func (q *queueLink) rPush(val string) {
	ele := &Element{
		val:  val,
		next: nil,
	}
	if q.empty() {
		q.head = ele
	} else {
		q.tail.next = ele
	}
	q.tail = ele
	q.length++
}

func (q *queueLink) lPop() (string, error) {
	if q.empty() {
		return "", ErrorEmpty
	}
	ele := q.head
	q.head = q.head.next
	q.length--
	if q.length == 0 {
		q.tail = nil
	}
	return ele.val, nil
}

func (q *queueLink) rPop() (string, error) {
	if q.empty() {
		return "", ErrorEmpty
	}
	q.length--
	temp := q.head
	if temp.next == nil {
		q.head = nil
		q.tail = nil
		return temp.val, nil
	}
	for temp.next.next != nil {
		temp = temp.next
	}
	ele := temp.next
	temp.next = nil
	q.tail = temp
	return ele.val, nil
}

func (q *queueLink) empty() bool {
	return q.length == 0
}
func (q *queueLink) Length() int {
	return q.length
}

func (q *queueLink) print() {
	fmt.Println("head:", q.head, "tail", q.tail, "length", q.length)
	temp := q.head
	for temp != nil {
		fmt.Printf("%s ", temp.val)
		temp = temp.next
	}
	fmt.Println()
}

// problem_1 回文字符串判断
func huiWen(str string) bool {
	stack := NewStackLink()
	for _, v := range str {
		stack.Push(string(v))
	}
	reverseStr := ""
	for {
		if v, err := stack.Pop(); err == nil {
			reverseStr += v
		} else {
			break
		}
	}
	return reverseStr == str
}

// problem_2 括号匹配判断
func checkParentheses(str string) bool {
	stack := NewStackLink()
	for _, v := range str {
		if v == '(' {
			stack.Push("(")
		}
		if v == ')' {
			_, err := stack.Pop()
			if err != nil {
				return false
			}
		}
	}
	return stack.empty()
}
