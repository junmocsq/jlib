package ch3

import (
	"strconv"
)

const (
	lparen  = '('
	rparen  = ')'
	plus    = '+'
	minus   = '-'
	times   = '*'
	divide  = '/'
	mod     = '%'
	eos     = ' '
	openand = 'O'
)

type expr struct {
	stackArr []int
	exprArr  []uint8
}

func NewExpression() *expr {
	return &expr{}
}

func (e *expr) getToken(c uint8) uint8 {
	if c >= '0' && c <= '9' {
		return openand
	}
	return c
}

// 后缀表达式求值
func (e *expr) eval(str string) (int, error) {
	numStr := ""
	for _, v := range str {
		symbol := e.getToken(uint8(v))
		if symbol == openand {
			numStr += string(uint8(v))
		} else {
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				e.stackPush(num)
				numStr = ""
			}
			if symbol == eos {
				continue
			}
			op2, _ := e.stackPop()
			op1, _ := e.stackPop()
			switch symbol {
			case plus:
				e.stackPush(op1 + op2)
			case minus:
				e.stackPush(op1 - op2)
			case times:
				e.stackPush(op1 * op2)
			case divide:
				e.stackPush(op1 / op2)
			case mod:
				e.stackPush(op1 % op2)
			}
		}
	}
	return e.stackPop()
}

// 正常表达式转换为后缀表达式，支持括号
func (e *expr) postfix(str string) string {
	var levelMap = make(map[uint8]int)
	levelMap[plus] = 1
	levelMap[minus] = 1
	levelMap[times] = 2
	levelMap[divide] = 2
	levelMap[mod] = 2
	levelMap[lparen] = 3
	levelMap[rparen] = 4
	str += " " // 添加空格保证最后的数字添加在后缀表达式里
	numStr := ""
	result := ""
	for _, v := range str {
		symbol := e.getToken(uint8(v))
		if symbol == openand {
			numStr += string(uint8(v))
		} else {
			if numStr != "" {
				result += " " + numStr
				numStr = ""
			}
			if symbol == eos {
				continue
			}
			if symbol == lparen { // 左括号 直接入栈
				e.exprPush(symbol)
			} else if symbol == rparen { // 右括号 出栈直到找到左括号为止
				for {
					op1, err := e.exprPop()
					if err != nil {
						break
					}
					if op1 == lparen {
						break
					}
					result += " " + string(op1)
				}
			} else { // 如果栈里的运算符优先级大于等于当前的运算符优先级，则出栈连接在表达式尾部
				op, err := e.exprTop()
				if err != nil {
					e.exprPush(symbol)
					continue
				}
				if op != lparen && levelMap[op] >= levelMap[symbol] {
					e.exprPop()
					result += " " + string(op)
				}
				e.exprPush(symbol)
			}
		}
	}
	for !e.exprEmpty() {
		v, _ := e.exprPop()
		result += " " + string(v)
	}
	return result
}

func (e *expr) stackPush(v int) {
	e.stackArr = append(e.stackArr, v)
}

func (e *expr) stackPop() (int, error) {
	if e.stackEmpty() {
		return 0, ErrorStackEmpty
	}
	length := len(e.stackArr)
	val := e.stackArr[length-1]
	e.stackArr = e.stackArr[:length-1]
	return val, nil
}

func (e *expr) stackEmpty() bool {
	return len(e.stackArr) == 0
}

func (e *expr) exprPush(v uint8) {
	e.exprArr = append(e.exprArr, v)
}

func (e *expr) exprTop() (uint8, error) {
	if e.exprEmpty() {
		return 0, ErrorStackEmpty
	}
	length := len(e.exprArr)
	return e.exprArr[length-1], nil
}

func (e *expr) exprPop() (uint8, error) {
	if e.exprEmpty() {
		return 0, ErrorStackEmpty
	}
	length := len(e.exprArr)
	val := e.exprArr[length-1]
	e.exprArr = e.exprArr[:length-1]
	return val, nil
}

func (e *expr) exprEmpty() bool {
	return len(e.exprArr) == 0
}
