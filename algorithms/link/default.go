package link

import (
	"reflect"
)

type singleNode struct {
	val  interface{}
	next *singleNode
}

type doubleNode struct {
	val  interface{}
	next *doubleNode
	pre  *doubleNode
}

type Linker interface {
	Find(val interface{}) int
	InsertByIndex(index int, val interface{}) bool
	ValueOf(index int) interface{}
	Add(values ...interface{}) bool
	Del(val interface{}) bool
	DelAll(val interface{}) int
	Empty() bool
	Length() int
	Elements() []interface{}
	Clear()
	Print()
}

func Equal(v1, v2 interface{}) bool {
	return reflect.DeepEqual(v1, v2)
}
