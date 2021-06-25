package ch3

import "errors"

type Element struct {
	Val int
}

var (
	ErrorStackEmpty = errors.New("栈为空")
	ErrorStackFull  = errors.New("栈已满")
	ErrorStacksFull = errors.New("栈已满")
	ErrorQueueEmpty = errors.New("队列为空")
	ErrorQueueFull  = errors.New("队列已满")
)

const (
	StackSize = 100
	QueueSize = 100
)
