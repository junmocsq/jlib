package ch6

import "testing"

func TestNewShortestPath(t *testing.T) {
	d := NewShortestPath(6)
	d.init()
	d.print()
	d.path(0)
}
