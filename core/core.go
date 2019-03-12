package core

import (
	"github.com/ddosakura/gRegEx/fa"
)

// NewFA create FA
func NewFA(fn func(s fa.State) fa.FA, expr []rune) (impl fa.FA, e error) {
	l := &Lexer{
		expr: expr,
		pos:  0,
	}
	p := &Parser{
		lexer: l,
	}
	return p.Start(fn)
}

// Match regexp
func Match(impl fa.FA, i fa.Inputer) bool {
	return fa.Verify(impl, i)
}
