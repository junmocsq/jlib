package algorithms

type Stacker interface {
	Push(interface{})
	Pop() interface{}
	Empty() bool
}
