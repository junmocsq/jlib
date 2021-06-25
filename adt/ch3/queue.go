package ch3

import (
	"fmt"
)

type Queue interface {
	IsFull() bool
	IsEmpty() bool
	PushRight(element *Element) error
	PushLeft(element *Element)
	PopLeft() (*Element, error)
	PopRight() (*Element, error)
}

type queue struct {
	arr         [QueueSize]*Element
	front, rear int
}

func NewQueue() *queue {
	return &queue{}
}
func (s *queue) IsFull() bool {
	return (s.front+QueueSize-1)%QueueSize == s.rear
}
func (s *queue) IsEmpty() bool {
	return s.front == s.rear
}
func (s *queue) nextRear() int {
	return (s.rear + 1) % QueueSize
}
func (s *queue) prevRear() int {
	return (s.rear - 1) % QueueSize
}
func (s *queue) nextFront() int {
	return (s.front + 1) % QueueSize
}
func (s *queue) prevFront() int {
	return (s.front - 1) % QueueSize
}
func (s *queue) PushRight(element *Element) error {
	if s.IsFull() {
		return ErrorQueueFull
	}
	s.arr[s.rear] = element
	s.rear = s.nextRear()
	return nil
}
func (s *queue) PushLeft(element *Element) error {
	if s.IsFull() {
		return ErrorQueueFull
	}
	s.front = s.prevFront()
	s.arr[s.front] = element
	return nil
}
func (s *queue) PopLeft() (*Element, error) {
	if s.IsEmpty() {
		return nil, ErrorStackEmpty
	}
	ele := s.arr[s.front]
	s.front = s.nextFront()
	return ele, nil
}
func (s *queue) PopRight() (*Element, error) {
	if s.IsEmpty() {
		return nil, ErrorStackEmpty
	}
	s.rear = s.prevRear()
	ele := s.arr[s.rear]
	return ele, nil
}

func (s *queue) print() {
	for i := s.front; i != s.rear; i = (i + 1) % QueueSize {
		fmt.Printf("%d=>%d \t", i, s.arr[i])
		//time.Sleep(time.Second)
	}
	fmt.Println()
}
