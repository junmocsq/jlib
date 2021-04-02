package algorithms

import "github.com/junmocsq/jlib/algorithms/link"

type Stacker interface {
	Push(interface{}) bool
	Pop() interface{}
	Empty() bool
}

const INIT_ARRAY_LENGTH = 10

type array struct {
	arr []interface{}
}

func NewArray() Stacker {
	return &array{
		arr: make([]interface{}, 0, INIT_ARRAY_LENGTH),
	}
}

func (a *array) Push(val interface{}) bool {
	a.arr = append(a.arr, val)
	return true
}

func (a *array) Pop() interface{} {
	if a.Empty() {
		return nil
	}
	size := len(a.arr)
	val := a.arr[size-1]
	a.arr = a.arr[:size-1]
	return val
}

func (a *array) Empty() bool {
	return len(a.arr) == 0
}

type stackLink struct {
	link link.Linker
}

func NewLink(linkTypes ...int) Stacker {
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
	return &stackLink{
		l,
	}
}

func (a *stackLink) Push(val interface{}) bool {
	return a.link.InsertByIndex(0, val)
}

func (a *stackLink) Pop() interface{} {
	return a.link.DelHead()
}

func (a *stackLink) Empty() bool {
	return a.link.Empty()
}
