package ch6

import "fmt"

type dijkstra struct {
	arr    [][]int
	length int
	MAX    int
}

func NewShortestPath(i int) *dijkstra {
	return &dijkstra{
		length: i,
		MAX:    1000,
	}
}

func (a *dijkstra) init() {
	a.arr = make([][]int, a.length)
	for k, _ := range a.arr {
		temp := make([]int, a.length)
		for _k, _ := range temp {
			if k == _k {
				temp[_k] = 0
			} else {
				temp[_k] = a.MAX
			}
		}
		a.arr[k] = temp
	}
	a.arr[0][1] = 50
	a.arr[0][2] = 10
	a.arr[0][4] = 45

	a.arr[1][2] = 15
	a.arr[1][4] = 10

	a.arr[2][0] = 20
	a.arr[2][3] = 15

	a.arr[3][1] = 20
	a.arr[3][4] = 35

	a.arr[4][3] = 30
	a.arr[5][3] = 3
}

func (d *dijkstra) print() {
	for _, v := range d.arr {
		for _, _v := range v {
			fmt.Printf("%4d\t", _v)
		}
		fmt.Println()
	}
}

// 单源最短路径算法
func (d *dijkstra) path(start int) {
	var path []int
	var distance, found []int
	distance = make([]int, d.length)
	found = make([]int, d.length)
	for i := 0; i < d.length; i++ {
		distance[i] = d.arr[start][i]
	}
	// 选择最小的边
	choose := func() int {
		min := d.MAX
		minpos := -1
		for i := 0; i < d.length; i++ {
			if distance[i] < min && found[i] == 0 {
				min = distance[i]
				minpos = i
			}
		}
		return minpos
	}
	found[start] = 1
	path = append(path, start)
	distance[start] = 0
	for i := 0; i < d.length-2; i++ {
		u := choose()
		found[u] = 1
		path = append(path, u)
		for w := 0; w < d.length; w++ {
			if found[w] == 0 {
				if distance[u]+d.arr[u][w] < distance[w] {
					distance[w] = distance[u] + d.arr[u][w]
				}
			}
		}
	}
	fmt.Println(distance, found, path)
}
