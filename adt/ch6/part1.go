package ch6

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

func (g *MatrixGraph) InsertVertex(v *Vertex) {

}

type Vertex struct {
	index int
	val   interface{}
}
