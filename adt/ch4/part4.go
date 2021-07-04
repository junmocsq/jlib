package ch4

import "fmt"

type polyNode struct {
	coef  int // 系数
	expon int // 幂
	link  *polyNode
}

type poly struct {
	head   *polyNode
	tail   *polyNode
	length int
}

func NewPoly() *poly {
	sentinel := &polyNode{}
	return &poly{
		head:   sentinel,
		tail:   sentinel,
		length: 0,
	}
}

func (p *poly) Clear() {
	p.head.link = nil
	p.tail = p.head
	p.length = 0
}

func (p *poly) Empty() bool {
	return p.length == 0
}

func (p *poly) Length() int {
	return p.length
}

func (p *poly) Insert(coef, expon int) {
	pre := p.FindPreByExpon(expon)
	ele := &polyNode{
		coef:  coef,
		expon: expon,
	}
	if pre.link == nil {
		pre.link = ele
		p.length++
	} else if pre.link.expon == expon {
		pre.link.coef = coef
	} else {
		ele.link = pre.link
		pre.link = ele
		p.length++
	}
}

func (p *poly) Add(coef, expon int) {
	pre := p.FindPreByExpon(expon)
	ele := &polyNode{
		coef:  coef,
		expon: expon,
	}
	p.AddWithPreAndNode(pre, ele)
}

func (p *poly) AddWithPreAndNode(pre *polyNode, ele *polyNode) {

	if pre.link == nil {
		pre.link = ele
		p.length++
	} else if pre.link.expon == ele.expon {
		pre.link.coef += ele.coef
	} else {
		ele.link = pre.link
		pre.link = ele
		p.length++
	}
}

func (p *poly) PAdd(p2 *poly) {
	pre1 := p.head
	node2 := p2.head.link
	p2Num := p2.length
	for pre1.link != nil && node2 != nil {
		if pre1.link.expon <= node2.expon {
			ele := node2
			node2 = node2.link
			ele.link = nil
			p.AddWithPreAndNode(pre1, ele)
			p2Num--

		}
		pre1 = pre1.link
	}
	if node2 != nil {
		pre1.link = node2
		p.tail = p2.tail
		p.length += p2Num
	}
	p2.Clear()
}

func (p *poly) FindPreByExpon(expon int) *polyNode {
	pre := p.head
	for pre.link != nil {
		if pre.link.expon <= expon {
			break
		}
		pre = pre.link
	}
	return pre
}

// 删除
func (p *poly) Earse(expon int) *polyNode {
	pre := p.FindPreByExpon(expon)
	if pre.link.expon == expon {
		node := pre.link
		pre.link = node.link
		p.length--
		if node.link == nil {
			p.tail = pre
		}
		return node
	}
	return nil
}

// 删除
func (p *poly) Print() string {
	str := ""
	pre := p.head.link
	for i := 0; i < p.length; i++ {
		if i == p.length-1 {
			str += fmt.Sprintf("%dx^%d", pre.coef, pre.expon)
		} else {
			str += fmt.Sprintf("%dx^%d + ", pre.coef, pre.expon)
		}
		pre = pre.link
	}
	return str
}
