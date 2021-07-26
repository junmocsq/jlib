package ch6

import "fmt"

// 图G由2个集合组成：一个顶点（vertex）构成的有穷非空集合和一个由边（edge）构成的有穷允空集合。V(G)和E(G)分别表示图G的顶点集和边集，有时也用G=(V,E)表示图。
// 无向图：是表示边的两个顶点之间没有次序关系的图。
// 顶点对(v0,v1)和(v1,v0)表示同一条边。
// 有向图：表示边的顶点对有方向的图。
// 顶点对<v0,v1>表示以顶点 v0 为尾(tail)，而以顶点 v1 为头(head) 的边。
// 完全图：是具有最多边数的图。有向图边数为 n*(n-1)；无向图边为 n*(n-1)/2
// 图的限制：
//		图中不能出现从顶点i到自身的边。
//		同一条边在图中不能出现两次或两次以上，不满足这个限制的图称为多重图(multigraph)
// 如果 (v0,v1) 是无向图的一条边，则称顶点v0和v1是相邻的,并称边(v0,v1)与顶点v0和v1关联
// 如果 <v0,v1> 是有向图的一条边，则称顶点 v0 邻接到顶点 v1，而 v1 邻接于 v0，并称边 <v0,v1>与顶点v0和v1关联。
// 路径（path）从顶点vp到顶点vq的一条路径是一个顶点序列vp，vi，...,Vq。路径的长度是路径上边的条数
// 一条`简单路径`是指路径上除了起点和终点可能相同外，其余顶点都互不相同的路径
// 图上的`回路`又称为`环路`是一条简单路径，且其起点和终点相同。
// 在无向图G中，如果从顶点v0到v1存在一条路径，则称`顶点v0和v1是连通的(connected)`。如果无向图G中每对顶点vi和vj，都存在一条从vi到vj的路径，则称`无向图G是连通的`。
// 无向图的连通分支，简称分支，是无向图中的极大连通子图。
// 在有向图G中，如果每对不同的顶点vi和vi，都存在 vi到vj 和 vj到vi 的有向路径，则称有向图G是强连通的。
// 顶点的度（degree）是关联于该顶点的边的条数。在有向图中，把顶点v的入度（in-degree）定义为以v为头（箭头指向v）的边的条数，而其出度（out-degree）定义以v为尾的边的条数。
// 无向图graph，有向图digraph

// 图的存储：邻接矩阵（adjacency matrices），邻接表（adjacency lists）和邻接多重表（adjacency multilists）

type grapher interface {
	InsertVertex(v *Vertex)    // 像图中插入没有关联边的新顶点v
	InsertEdge(v1, v2 *Vertex) // 向图的顶点v1和v2之间插入一条边
	DeleteVertex(v *Vertex)    // 删除图中的顶点v及其关联的所有边
	DeleteEdge(v1, v2 *Vertex) // 删除图中边(v1,v2),顶点v1，v2不删除
	IsEmpty() bool
	Adjacent(v *Vertex) []*Vertex // 顶点v的所有邻接节点
}

type MatrixGraph struct {
	arr    [][]bool
	vertex []*Vertex
}

func CreateMatrixGraph(vertexNum int) *MatrixGraph {
	arr := make([][]bool, vertexNum)
	for k, _ := range arr {
		arr[k] = make([]bool, vertexNum)
	}
	return &MatrixGraph{
		arr:    arr,
		vertex: make([]*Vertex, vertexNum),
	}
}

type Vertex struct {
	index int
	next  *Vertex
}

type AdjacencyLists struct {
	vertexes []*Vertex
}

func NewAdjacencyLists(size int) *AdjacencyLists {
	return &AdjacencyLists{
		vertexes: make([]*Vertex, size),
	}
}

// InsertVertex 向图中插入没有关联边的新顶点v
func (l *AdjacencyLists) InsertVertex(v *Vertex) {
	l.vertexes[v.index] = v
}

// DeleteVertex 删除图中的顶点v及其关联的所有边
func (l *AdjacencyLists) DeleteVertex(v *Vertex) {
	list := l.vertexes[v.index]
	for list.next != nil {
		temp := l.vertexes[list.next.index]
		for temp.next != nil {
			if temp.next.index == v.index {
				temp.next = temp.next.next
			} else {
				temp = temp.next
			}
		}
		list = list.next
	}
	l.vertexes[v.index] = nil
}

// DeleteEdge 删除图中边(v1,v2),顶点v1，v2不删除
func (l *AdjacencyLists) DeleteEdge(v1, v2 *Vertex) {
	temp1 := l.vertexes[v1.index]
	for temp1.next != nil {
		if temp1.next.index == v2.index {
			temp1.next = temp1.next.next
		} else {
			temp1 = temp1.next
		}
	}
	temp2 := l.vertexes[v2.index]
	for temp2.next != nil {
		if temp2.next.index == v1.index {
			temp2.next = temp2.next.next
		} else {
			temp2 = temp2.next
		}
	}
}

func (l *AdjacencyLists) IsEmpty() bool {
	result := true
	for _, v := range l.vertexes {
		if v != nil {
			result = false
			break
		}
	}
	return result
}

// Adjacent 顶点v的所有邻接节点
func (l *AdjacencyLists) Adjacent(v *Vertex) []*Vertex {
	return l.vertexes
}

func (l *AdjacencyLists) Print() {
	fmt.Println("--------------------------------")
	for k, v := range l.vertexes {
		fmt.Printf("%d: ", k)
		for v != nil {
			if v.next != nil {
				fmt.Printf("%d=>", v.index)
			} else {
				fmt.Printf("%d", v.index)
			}
			v = v.next
		}
		fmt.Println()
	}
	fmt.Println("--------------------------------")
}

// InsertEdge 向图的顶点v1和v2之间插入一条边
func (l *AdjacencyLists) InsertEdge(v1, v2 *Vertex) {
	temp1 := l.vertexes[v1.index]
	if temp1 == nil {
		l.vertexes[v1.index] = &Vertex{index: v1.index}
		temp1 = l.vertexes[v1.index]
	}
	for temp1.next != nil {
		if temp1.next.index == v2.index { // 不添加重复节点
			return
		}
		temp1 = temp1.next
	}
	temp1.next = v2

	temp2 := l.vertexes[v2.index]
	if temp2 == nil {
		l.vertexes[v2.index] = &Vertex{index: v2.index}
		temp2 = l.vertexes[v2.index]
	}
	for temp2.next != nil {
		if temp2.next.index == v1.index { // 不添加重复节点
			return
		}
		temp2 = temp2.next
	}
	temp2.next = v1
}

func (l *AdjacencyLists) dfsRecursion() []int {
	var fRecv func(v *Vertex)
	visited := make([]bool, len(l.vertexes))
	var res []int
	fRecv = func(v *Vertex) {
		if !visited[v.index] {
			res = append(res, v.index)
			visited[v.index] = true
		}
		for v.next != nil {
			if !visited[v.next.index] {
				fRecv(l.vertexes[v.next.index])
			}
			v = v.next
		}
	}
	fRecv(l.vertexes[3])
	return res
}

func (l *AdjacencyLists) dfs() []int {
	var result []int
	var fStack func(index int)
	s := stack()
	fStack = func(index int) {
		visited := make([]bool, len(l.vertexes))
		s.stackPush(l.vertexes[index])
		for !s.stackEmpty() {
			temp, _ := s.stackPop() // 顶点w
			temp1 := temp.next
			for temp1 != nil { // 顶点w所在链表的下一个未访问节点入栈
				if !visited[temp1.index] {
					s.stackPush(temp1)
					break
				}
				temp1 = temp1.next
			}
			if !visited[temp.index] {
				result = append(result, temp.index)
				visited[temp.index] = true
				v := l.vertexes[temp.index].next

				for v != nil { // 顶点w的头结点的下一个未访问节点入栈
					if !visited[v.index] {
						s.stackPush(v)
						break
					}
					v = v.next
				}
			}
		}
	}
	fStack(0)
	return result
}

func (l *AdjacencyLists) dfsTree() [][2]int {
	var result [][2]int
	var fStack func(index int)
	s := stack()
	fStack = func(index int) {
		visited := make([]bool, len(l.vertexes))
		s.stackPush(l.vertexes[index]) // head
		s.stackPush(l.vertexes[index]) // node
		for !s.stackEmpty() {
			temp, _ := s.stackPop() // 顶点w
			head, _ := s.stackPop() // 顶点w
			temp1 := temp.next
			for temp1 != nil { // 顶点w所在链表的下一个未访问节点入栈
				if !visited[temp1.index] {
					s.stackPush(head)
					s.stackPush(temp1)
					break
				}
				temp1 = temp1.next
			}
			if !visited[temp.index] {
				if head.index != temp.index {
					result = append(result, [2]int{head.index, temp.index})
				}
				visited[temp.index] = true
				head = l.vertexes[temp.index]
				v := head.next
				for v != nil { // 顶点w的头结点的下一个未访问节点入栈
					if !visited[v.index] {
						s.stackPush(head)
						s.stackPush(v)
						break
					}
					v = v.next
				}
			}
		}
	}
	fStack(0)
	return result
}

// 深度优先生成树
func (l *AdjacencyLists) dfsRecursionTree() [][2]int {
	var fRecv func(v *Vertex)
	visited := make([]bool, len(l.vertexes))
	var edgeRes [][2]int
	fRecv = func(v *Vertex) {
		if !visited[v.index] {
			visited[v.index] = true
		}
		head := v
		for v.next != nil {
			if !visited[v.next.index] {
				edgeRes = append(edgeRes, [2]int{head.index, v.next.index})
				visited[v.next.index] = true
				fRecv(l.vertexes[v.next.index])
			}
			v = v.next
		}
	}
	fRecv(l.vertexes[0])
	return edgeRes
}
func (l *AdjacencyLists) dfnlow() {
	dfn := make([]int, len(l.vertexes))
	low := make([]int, len(l.vertexes))
	//visited := make([]bool,len(l.vertexes))
	for k, _ := range l.vertexes {
		dfn[k] = -1
		low[k] = -1
	}
	var f func(u, v int)
	num := 0
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	var res []int
	f = func(u, v int) {

		res = append(res, u)
		dfn[u] = num
		low[u] = num
		num++
		head := l.vertexes[u]
		for head != nil {
			w := head.index

			if dfn[w] < 0 {
				f(w, u)
				if u==1{
					//fmt.Println(v,u,w,low,dfn)
				}
				low[u] = min(low[u], low[w])
			} else if w != v {
				if u==1{
					//fmt.Println("--",v,u,w,low,dfn)
				}
				low[u] = min(low[u], dfn[w])
			}
			head = head.next
		}
	}
	f(3, -1)
	fmt.Println(res, dfn, low)
}

func (l *AdjacencyLists) bfs() []int {
	var result []int
	var fQueue func(index int)
	q := queue()
	fQueue = func(index int) {
		visited := make([]bool, len(l.vertexes))
		q.addq(l.vertexes[index])
		for !q.emptyq() {
			temp, _ := q.deleteq()
			for temp != nil {
				if !visited[temp.index] {
					result = append(result, temp.index)
					visited[temp.index] = true
					q.addq(l.vertexes[temp.index])
				}
				temp = temp.next
			}
		}
	}
	fQueue(0)
	return result
}

// 广度优先生成树
func (l *AdjacencyLists) bfsTree() [][2]int {
	var result [][2]int
	var fQueue func(index int)
	q := queue()
	fQueue = func(index int) {
		visited := make([]bool, len(l.vertexes))
		q.addq(l.vertexes[index])
		for !q.emptyq() {
			temp, _ := q.deleteq()
			head := temp
			for temp != nil {
				if !visited[temp.index] {
					if head.index != temp.index {
						result = append(result, [2]int{head.index, temp.index})
					}
					visited[temp.index] = true
					q.addq(l.vertexes[temp.index])
				}
				temp = temp.next
			}
		}
	}
	fQueue(0)
	return result
}

type stackArr struct {
	stackArr []*Vertex
}

func stack() *stackArr {
	return &stackArr{}
}
func (l *stackArr) stackEmpty() bool {
	return len(l.stackArr) == 0
}
func (l *stackArr) stackClear() {
	l.stackArr = nil
}
func (l *stackArr) stackPush(v *Vertex) {
	l.stackArr = append(l.stackArr, v)
}
func (l *stackArr) stackPop() (v *Vertex, ok bool) {
	if l.stackEmpty() {
		return
	}
	v = l.stackArr[len(l.stackArr)-1]
	l.stackArr = l.stackArr[:len(l.stackArr)-1]
	return v, true
}

type queueArr struct {
	queueArr []*Vertex
}

func queue() *queueArr {
	return &queueArr{}
}
func (l *queueArr) deleteq() (v *Vertex, ok bool) {
	if l.emptyq() {
		return
	}
	v = l.queueArr[0]
	l.queueArr = l.queueArr[1:]
	return v, true
}
func (l *queueArr) addq(v *Vertex) {
	l.queueArr = append(l.queueArr, v)
}
func (l *queueArr) emptyq() bool {
	return len(l.queueArr) == 0
}
func (l *queueArr) clearq() {
	l.queueArr = nil
}
