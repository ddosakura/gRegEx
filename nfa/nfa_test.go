package nfa

import (
	"testing"

	"github.com/ddosakura/gRegEx/fa"
)

func TestBasic(t *testing.T) {
	n0 := fa.NewBasicState(false)
	n1 := fa.NewBasicState(false)
	n2 := fa.NewBasicState(true)
	a := fa.NewRuneInput('a')
	b := fa.NewRuneInput('b')

	nfa := New(n0)
	nfa.AddState(n1)
	nfa.AddState(n2)

	nfa.AddTransition(n0, a, n1)
	nfa.AddTransition(n1, b, n1, n2)

	if ret := nfa.Next(a); len(ret) < 1 || ret[0] != n1 {
		t.Errorf("Expect %p, but get %+v\n", n1, ret)
	} else {
		t.Logf("%+v\n", ret)
	}

	if ret := nfa.Next(b); len(ret) < 1 || ret[0] != n1 || ret[1] != n2 {
		t.Errorf("Expect [%p %p], but get %+v\n", n1, n2, ret)
	} else {
		t.Logf("%+v\n", ret)
	}

	if ret := nfa.Next(b); len(ret) < 1 || ret[0] != n1 || ret[1] != n2 {
		t.Errorf("Expect [%p %p], but get %+v\n", n1, n2, ret)
	} else {
		t.Logf("%+v\n", ret)
	}

	if !nfa.IsFinal() {
		t.Errorf("Verify is failed")
	}
}
