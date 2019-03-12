package core

import (
	"errors"
	"strconv"
)

// Lexer of regexp
type Lexer struct {
	expr []rune
	pos  int
}

// Next read rune
func (l *Lexer) Next() (s *Symbol, e error) {
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		pos := e.(int)
		e = errors.New("lexer error in " + strconv.Itoa(pos))
	}()
	if l.pos == len(l.expr) {
		return nil, nil
	}
	r := l.expr[l.pos]
	switch r {
	case '\\':
		l.pos++
		if l.pos == len(l.expr) {
			panic(l.pos)
		}
		s = token[l.expr[l.pos]]
		if s == nil {
			panic(l.pos)
		}
	case '^':
		s = SOL
	case '$':
		s = EOL
	case '[':
		s = NewSymbol()
		l.pos++
		if l.pos == len(l.expr) {
			panic(l.pos)
		}
		if l.expr[l.pos] == '^' {
			s.Opposite = true
		} else {
			l.pos--
		}

		for l.expr[l.pos] != ']' {
			l.pos++
			if l.pos == len(l.expr) {
				panic(l.pos)
			}
			lr := l.expr[l.pos]
			if lr == '\\' {
				l.pos++
				if l.pos == len(l.expr) {
					panic(l.pos)
				}
				lr = l.expr[l.pos]
			}

			l.pos++
			if l.pos == len(l.expr) {
				panic(l.pos)
			}
			if l.expr[l.pos] == '-' {
				l.pos++
				if l.pos == len(l.expr) {
					panic(l.pos)
				}
				rr := l.expr[l.pos]
				if rr == '\\' {
					l.pos++
					if l.pos == len(l.expr) {
						panic(l.pos)
					}
					rr = l.expr[l.pos]
				}
				for i := lr; i <= rr; i++ {
					s.Raw[i] = true
				}
			} else {
				l.pos--
				s.Raw[lr] = true
			}
		}

	case '|', '*', '+', '?':
		s = NewOp(r)
	default:
		s = NewSymbol(r)
	}
	l.pos++
	return
}
