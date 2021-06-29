package ch3

import (
	"fmt"
	"math/rand"
	"time"
)

type point struct {
	x, y int
}

// 迷宫 p72
type mg struct {
	wide    int
	height  int
	arr     [][]int
	visited [][]int
	stack   []point
}

func CreateMg(wide, height int) *mg {
	var arr = make([][]int, height)
	var visited = make([][]int, height)
	for k, _ := range arr {
		temp := make([]int, wide)
		for _k, _ := range temp {
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(2) >= 1 {
				temp[_k] = 1
			} else {
				temp[_k] = 0
			}
		}
		arr[k] = temp
		visited[k] = make([]int, wide)
	}
	return &mg{
		wide:    wide,
		height:  height,
		arr:     arr,
		visited: visited,
	}
}

func (m *mg) move(i int) (x int, y int) {
	switch i {
	case 0:
		return 0, -1
	case 1:
		return 1, -1
	case 2:
		return 1, 0
	case 3:
		return 1, 1
	case 4:
		return 0, 1
	case 5:
		return -1, 1
	case 6:
		return -1, 0
	case 7:
		return -1, -1
	}
	return 0, 0
}

func (m *mg) search() bool {
	m.push(0, 0)
	m.visited[0][0] = 1

	for !m.empty() {
		flag := false
		currX, currY := m.top()
		for i := 0; i < 8; i++ {
			moveX, moveY := m.move(i)
			nextX := currX + moveX
			nextY := currY + moveY
			if nextX < 0 || nextX > m.wide-1 {
				continue
			}
			if nextY < 0 || nextY > m.height-1 {
				continue
			}
			if nextX == m.wide-1 && nextY == m.height-1 {
				return true
			}
			if m.visited[nextY][nextX] == 1 {
				continue
			}
			flag = true
			m.visited[nextY][nextX] = 1
			if m.arr[nextY][nextX] == 1 {
				m.push(nextX, nextY)
				break
			}
		}
		if !flag {
			m.pop()
		}
	}
	return false
}

func (m *mg) print() {
	fmt.Println(m.stack)
	for !m.empty() {
		x, y := m.pop()
		m.arr[y][x] = 9
	}
	for k, v := range m.arr {
		fmt.Println(k, v)
	}
}

func (m *mg) push(x, y int) {
	m.stack = append(m.stack, point{
		x: x,
		y: y,
	})
}
func (m *mg) pop() (x, y int) {
	res := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	return res.x, res.y
}
func (m *mg) top() (x, y int) {
	res := m.stack[len(m.stack)-1]
	return res.x, res.y
}
func (m *mg) empty() bool {
	return len(m.stack) == 0
}
