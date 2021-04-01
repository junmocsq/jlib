package algorithms

import "github.com/junmocsq/jlib/algorithms/link"

type Queue interface {
	Enqueue(interface{}) bool
	Dequeue() interface{}
	Empty() bool
}

var _ Queue = &array{}

func NewArrayQueue() Queue {
	return &array{
		arr: make([]interface{}, 0, INIT_ARRAY_LENGTH),
	}
}

func (a *array) Enqueue(val interface{}) bool {
	a.arr = append(a.arr, val)
	return true
}

func (a *array) Dequeue() interface{} {
	if a.Empty() {
		return nil
	}
	val := a.arr[0]
	a.arr = a.arr[1:]
	return val
}

type queueLink struct {
	link link.Linker
}

func NewQueueLink(linkTypes ...int) Queue {
	linkType := 1
	if len(linkTypes) > 0 {
		linkType = linkTypes[0]
	}
	var l link.Linker
	switch linkType {
	case 1:
		l = link.NewSingle()
	case 2:
		l = link.NewCircular()
	case 3:
		l = link.NewDouble()
	case 4:
		l = link.NewDoubleCircular()
	default:
		l = link.NewSingle()
	}
	return &queueLink{
		l,
	}
}

func (a *queueLink) Enqueue(val interface{}) bool {
	return a.link.InsertByIndex(0, val)
}

func (a *queueLink) Dequeue() interface{} {
	return a.link.DelTail()
}

func (a *queueLink) Empty() bool {
	return a.link.Empty()
}

type loopArray struct {
	arr      *[INIT_ARRAY_LENGTH]interface{}
	capacity int
	head     int
	tail     int
}

// 循环数组实现
func NewLoopArrayStack() Stacker {
	return &loopArray{
		arr:      &[INIT_ARRAY_LENGTH]interface{}{},
		capacity: INIT_ARRAY_LENGTH,
	}
}

// 循环数组实现
func NewLoopArrayQueue() Queue {
	return &loopArray{
		arr:      &[INIT_ARRAY_LENGTH]interface{}{},
		capacity: INIT_ARRAY_LENGTH,
	}
}

func (l *loopArray) Push(val interface{}) bool {
	if l.Full() {
		return false
	}
	l.arr[l.tail] = val
	l.tail = l.next(l.tail)
	return true
}

func (l *loopArray) size() int {
	return (l.tail + l.capacity - l.head) % l.capacity
}

func (l *loopArray) next(index int) int {
	return (index + 1 + l.capacity) % l.capacity
}

func (l *loopArray) pre(index int) int {
	return (index - 1 + l.capacity) % l.capacity
}

func (l *loopArray) Full() bool {
	if l.size() < l.capacity-1 {
		return false
	}
	return true
}

func (l *loopArray) Pop() interface{} {
	if l.Empty() {
		return nil
	}
	l.tail = l.pre(l.tail)
	return l.arr[l.tail]
}

func (l *loopArray) Enqueue(val interface{}) bool {
	if l.Full() {
		return false
	}
	l.arr[l.tail] = val
	l.tail = l.next(l.tail)
	return true
}

func (l *loopArray) Dequeue() interface{} {
	if l.Empty() {
		return nil
	}
	pre := l.head
	l.head = l.next(l.head)
	return l.arr[pre]
}

func (l *loopArray) Empty() bool {
	return l.size() == 0
}
