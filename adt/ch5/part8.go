package ch5

// TODO 败者树 没有数据子树处理异常
/*

胜者树：每个父结点表示它的两个子女中比赛胜利的结点，胜利的结点继续向上比赛。根结点记录了最后的胜者。

           重构：从缓冲区中再取出一个结点，将其与自己的兄弟结点比较，如果胜，则继续向上比较，否则，其兄弟代替它继续向上比较。

败者树：每个父结点表示它的两个子女中比赛失败的结点，胜利的结点继续向上比赛。根结点记录的是与最后的胜者比较后失败者，所以需要添加 一个结点来表示最后的胜者。

           重构：从缓冲区中再取出一个结点，将其与自己的父亲结点进行比较，如果胜，则继续向上比较，否则，其父亲代替它继续向上比较。

           注：叶子结点记录的是真实的值，而非叶结点记录了相应的结点标号，不是实际的数据

                 胜者树和败者树的复杂度相同，只是败者树的重构过程比较简单
*/
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type selectNode struct {
	val   int
	index int
}
type selectTree struct {
	arr       []*selectNode   // 选择树
	list      [][]*selectNode // 比较的有序队列
	leafStart int             // 叶节点起点
}

func NewSelectTree() *selectTree {
	return &selectTree{}
}

// 创造测试数据 有序队列数组
func createTestSortListArr(bucketNum int) [][]*selectNode {
	var arr []int
	for i := 0; i < 200; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(1000))
	}
	sort.Ints(arr)
	m := map[int]int{}
	for i := 1; i < len(arr); i++ {
		m[arr[i]] = arr[i]
	}
	res := make([][]int, bucketNum)
	index := 0
	for k := range m {
		res[index%bucketNum] = append(res[index%bucketNum], k)
		index++
	}
	for k := range res {
		sort.Ints(res[k])
	}
	result := make([][]*selectNode, bucketNum)
	for k := range res {
		var temp []*selectNode
		for _, v := range res[k] {
			temp = append(temp, &selectNode{
				index: k,
				val:   v,
			})
		}
		result[k] = temp
	}
	return result
}

func (s *selectTree) min(a int, b int) int {
	if s.arr[a] == nil && s.arr[b] == nil {
		return -1
	}
	if s.arr[a] == nil {
		return b
	}
	if s.arr[b] == nil {
		return a
	}
	if s.arr[a].val < s.arr[b].val {
		return a
	} else {
		return b
	}
}

// min max
func (s *selectTree) minMax(a, b *selectNode) (*selectNode, *selectNode) {
	if a != nil && b != nil {
		if a.val < b.val {
			return a, b
		} else {
			return b, a
		}

	}
	return a, b

}

func (s *selectTree) arrLength(leaf int) int {
	res := 1
	for {
		if res >= leaf {
			return res
		}
		res = res * 2
	}
}

func (s *selectTree) parent(index int) int {
	return index / 2
}

func (s *selectTree) left(index int) int {
	return index * 2
}
func (s *selectTree) right(index int) int {
	return index*2 + 1
}

func (s *selectTree) CreateSuccessTree(bucketNum int) {
	res := createTestSortListArr(bucketNum)
	leafLength := s.arrLength(len(res))
	s.leafStart = leafLength
	length := 2 * leafLength // 数组长度为叶节点的2倍，对于胜者树0号不存数据
	s.arr = make([]*selectNode, length)
	// 初始化叶子节点和队列
	for k, v := range res {
		if len(v) > 0 {
			s.arr[leafLength+k] = v[0]
			s.list = append(s.list, v[1:])
		}
	}
	// 初始化非叶子结点
	for total := leafLength / 2; total > 0; total /= 2 {
		for k := total; k < 2*total; k++ {
			left := s.left(k)
			minIndex := s.min(left, left+1)
			if minIndex > 0 {
				s.arr[k] = s.arr[minIndex]
			}

		}
	}

	//for k, v := range s.arr {
	//	if v != nil {
	//		fmt.Printf("%d:%d(%d) ", k, v.val, v.index)
	//	}
	//}
	//fmt.Println()

}

func (s *selectTree) SortSuccessTree() []int {
	var arr []int
	for s.arr[1] != nil {
		temp := s.arr[1]
		arr = append(arr, temp.val)
		//time.Sleep(time.Second)
		index := temp.index
		nodeIndex := s.leafStart + index
		if len(s.list[index]) == 0 {
			s.arr[nodeIndex] = nil
		} else {
			s.arr[nodeIndex] = s.list[index][0]
			s.list[index] = s.list[index][1:]
		}
		parent := s.parent(nodeIndex)
		for parent > 0 {
			nLeft := s.left(parent)
			minIndex := s.min(nLeft, nLeft+1)
			if minIndex < 0 {
				s.arr[parent] = nil
			} else {
				s.arr[parent] = s.arr[minIndex]
			}
			parent = s.parent(parent)
		}
	}
	fmt.Println(s.arr)
	return arr
}

func (s *selectTree) CreateFailedTree(bucketNum int) {
	res := createTestSortListArr(bucketNum)
	leafLength := s.arrLength(len(res))
	s.leafStart = leafLength
	length := 2 * leafLength // 数组长度为叶节点的2倍，对于胜者树0号不存数据
	s.arr = make([]*selectNode, length)
	// 初始化叶子节点和队列
	tempSuccessArr := make([]*selectNode, length)
	for k, v := range res {
		if len(v) > 0 {
			s.arr[leafLength+k] = v[0]
			tempSuccessArr[leafLength+k] = v[0]
			s.list = append(s.list, v[1:])
		}
	}
	// 初始化非叶子结点
	for total := leafLength / 2; total > 0; total /= 2 {
		fmt.Println("total", total)
		for k := total; k < 2*total; k++ {
			left := s.left(k)
			minNode, maxNode := s.minMax(tempSuccessArr[left], tempSuccessArr[left+1])
			s.arr[k] = maxNode
			tempSuccessArr[k] = minNode
			s.arr[0] = minNode // 当前最小值存0
		}
	}

	for _, v := range s.arr {
		if v != nil {
			fmt.Printf("%d ", v.val)
		}
	}
	fmt.Println()

}

func (s *selectTree) SortFailedTree() []int {
	var arr []int
	for {
		index := -1
		if s.arr[0] != nil {
			arr = append(arr, s.arr[0].val)
			index = s.arr[0].index
			s.arr[0] = nil
		} else {
			for k := 1; k < len(s.arr); k++ {
				if s.arr[k] != nil {
					index = s.arr[k].index
					break
				}
			}
		}

		if index == -1 {
			break
		}

		nodeIndex := s.leafStart + index

		if len(s.list[index]) == 0 {
			s.arr[nodeIndex] = nil
			for k := 1; k < len(s.arr); k++ {
				if s.arr[k] != nil {
					index = s.arr[k].index
					break
				}
			}
			if index == -1 {
				break
			}
			nodeIndex = s.leafStart + index
			if s.arr[nodeIndex] == nil {
				fmt.Println(s.arr, nodeIndex)
				break
			}
		} else {
			s.arr[nodeIndex] = s.list[index][0]
			s.list[index] = s.list[index][1:]
		}

		succNode := s.arr[nodeIndex]
		parent := s.parent(nodeIndex)
		fmt.Println(index, s.arr[nodeIndex], parent, s.list[index])
		time.Sleep(10 * time.Millisecond)
		if succNode == nil && s.arr[parent] == nil {
			continue
		}

		for parent > 0 {
			//if succNode == nil {
			//	succNode = s.arr[parent]
			//	s.arr[parent] = nil
			//} else {
			//	if s.arr[parent] != nil && s.arr[parent].val < succNode.val {
			//		succNode = s.arr[parent]
			//		s.arr[parent] = s.arr[nodeIndex]
			//	}
			//}
			if succNode == s.arr[parent] {
				s.arr[parent] = nil
			}
			if s.arr[parent] != nil && s.arr[parent].val < succNode.val {
				succNode = s.arr[parent]
				s.arr[parent] = s.arr[nodeIndex]
			}
			nodeIndex = parent
			parent = s.parent(parent)
		}
		s.arr[0] = succNode
	}
	fmt.Println(s.arr)
	return arr
}
