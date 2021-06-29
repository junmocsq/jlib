package ch3

import "testing"

func TestNewExpression(t *testing.T) {
	s := "600 2/3-400 200* +"
	ex := NewExpression()
	wanted := 80297
	if actual, e := ex.eval(s); e != nil {
		if actual != wanted {
			t.Errorf("%s wanted:%d ,actual:%d", s, wanted, actual)
		}
	}

	s = "600 2/100-400 200* +"
	ex = NewExpression()
	wanted = 80200
	if actual, e := ex.eval(s); e != nil {
		if actual != wanted {
			t.Errorf("%s wanted:%d ,actual:%d", s, wanted, actual)
		}
	}

	s = "6/2-3+4*2"
	nn := NewExpression()
	exprStr := nn.postfix(s)
	wanted = 8
	if actual, e := ex.eval(exprStr); e != nil {
		if actual != wanted {
			t.Errorf("%s wanted:%d ,actual:%d", s, wanted, actual)
		}
	}

}
