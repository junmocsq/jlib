package ch6

import "testing"

func TestAdjacencyLists_Adjacent(t *testing.T) {
	l := NewAdjacencyLists(10)
	l.InsertEdge(&Vertex{index: 0}, &Vertex{index: 1})
	l.InsertEdge(&Vertex{index: 0}, &Vertex{index: 2})
	l.InsertEdge(&Vertex{index: 1}, &Vertex{index: 3})
	l.InsertEdge(&Vertex{index: 1}, &Vertex{index: 4})
	l.InsertEdge(&Vertex{index: 2}, &Vertex{index: 5})
	l.InsertEdge(&Vertex{index: 2}, &Vertex{index: 6})
	l.InsertEdge(&Vertex{index: 3}, &Vertex{index: 7})
	l.InsertEdge(&Vertex{index: 4}, &Vertex{index: 7})
	l.InsertEdge(&Vertex{index: 5}, &Vertex{index: 7})
	l.InsertEdge(&Vertex{index: 6}, &Vertex{index: 7})
	l.InsertEdge(&Vertex{index: 6}, &Vertex{index: 7})
	l.InsertEdge(&Vertex{index: 0}, &Vertex{index: 8})
	l.InsertEdge(&Vertex{index: 0}, &Vertex{index: 9})

	l.Print()
	t.Log(l.dfs())
	t.Log(l.dfsRecursion())
	t.Log(l.dfsTree())
	t.Log(l.dfsRecursionTree())
	t.Log(l.bfs())

	t.Log(l.bfsTree())



	l = NewAdjacencyLists(10)
	l.InsertEdge(&Vertex{index: 0},&Vertex{index: 1})
	l.InsertEdge(&Vertex{index: 1},&Vertex{index: 2})
	l.InsertEdge(&Vertex{index: 3},&Vertex{index: 4})
	l.InsertEdge(&Vertex{index: 1},&Vertex{index: 3})
	l.InsertEdge(&Vertex{index: 2},&Vertex{index: 4})
	l.InsertEdge(&Vertex{index: 3},&Vertex{index: 5})
	l.InsertEdge(&Vertex{index: 5},&Vertex{index: 6})
	l.InsertEdge(&Vertex{index: 5},&Vertex{index: 7})
	l.InsertEdge(&Vertex{index: 6},&Vertex{index: 7})
	l.InsertEdge(&Vertex{index: 7},&Vertex{index: 9})
	l.InsertEdge(&Vertex{index: 7},&Vertex{index: 8})
	l.Print()
	t.Log(l.dfsRecursion())
	l.dfnlow()

}
