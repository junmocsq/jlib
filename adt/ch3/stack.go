package ch3

import "fmt"

type Stacker interface {
	IsFull() bool
	IsEmpty() bool
	Add(element *Element) error
	Delete() (*Element, error)
}

type stack struct {
	arr  [StackSize]*Element
	size int
}

func NewStack() *stack {
	return &stack{}
}
func (s *stack) IsFull() bool {
	return s.size == StackSize
}
func (s *stack) IsEmpty() bool {
	return s.size == 0
}
func (s *stack) Add(element *Element) error {
	if s.IsFull() {
		return ErrorStackFull
	}
	s.arr[s.size] = element
	s.size++
	return nil
}
func (s *stack) Delete() (*Element, error) {
	if s.IsEmpty() {
		return nil, ErrorStackEmpty
	}
	ele := s.arr[s.size-1]
	s.size--
	return ele, nil
}

func (s *stack) Copy() *stack {
	ns := NewStack()
	ns.size = s.size
	for i := 0; i < s.size; i++ {
		ns.arr[i] = s.arr[i]
	}
	return ns
}

//
func p1_4(n int) {
	//var result [][]int

	var f func(index int, s *stack, temp []int)
	f = func(index int, s *stack, temp []int) {
		if index == n {
			for !s.IsEmpty() {
				v, _ := s.Delete()
				temp = append(temp, v.Val)
			}
			fmt.Println("==", temp)
			return
		}
		s2 := s.Copy()

		s.Add(&Element{Val: index})
		t1 := make([]int, len(temp))
		copy(t1, temp)
		f(index+1, s, t1)

		temp = append(temp, index)
		t2 := make([]int, len(temp))
		copy(t2, temp)
		f(index+1, s2, t2)
	}

	f(0, NewStack(), []int{})
}
