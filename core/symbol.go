package core

import (
	"github.com/ddosakura/gRegEx/fa"
)

// SymbolType of regexp
type SymbolType int

const (
	_ SymbolType = iota
	// SymbolRaw is a rune in set
	SymbolRaw
	// SymbolAny is a rune in set with any rune
	SymbolAny
	// SymbolSOL is `Start of Line`
	SymbolSOL
	// SymbolEOL is `End of Line`
	SymbolEOL

	// SymbolOp for Symbol
	SymbolOp
)

// Symbol of regexp
type Symbol struct {
	Type     SymbolType
	Raw      map[rune]bool
	Opposite bool // only useful for SymbolRaw
	Op       rune // only useful for SymbolOp
}

// Eq for Symbol
func (s *Symbol) Eq(i fa.Input) (b bool) {
	if s.Type == SymbolAny {
		return true
	}
	defer func() {
		e := recover()
		if e != nil {
			b = false
		}
	}()
	// TODO: rewrite RuneInput
	o := i.(*fa.RuneInput)
	if s.Type == SymbolRaw {
		for r, p1 := range s.Raw {
			p2 := fa.NewRuneInput(r).Eq(o)
			if p2 {
				return p1 && !s.Opposite && p2
			}
		}
	}
	return s.Opposite
}

// NewSymbol of regexp
func NewSymbol(raw ...rune) (s *Symbol) {
	s = &Symbol{
		Type: SymbolRaw,
		Raw:  make(map[rune]bool),
	}
	for _, r := range raw {
		s.Raw[r] = true
	}
	return
}

// NewSymbolR reverse Symbol
func NewSymbolR(raw ...rune) (s *Symbol) {
	s = NewSymbol(raw...)
	s.Opposite = true
	return
}

// NewSymbolRange rangeSymbol
func NewSymbolRange(raw ...[]rune) (s *Symbol) {
	s = &Symbol{
		Type: SymbolRaw,
		Raw:  make(map[rune]bool),
	}
	for _, r := range raw {
		for i := r[0]; i <= r[1]; i++ {
			s.Raw[i] = true
		}
	}
	return
}

// NewSymbolRangeR reverse rangeSymbol
func NewSymbolRangeR(raw ...[]rune) (s *Symbol) {
	s = NewSymbolRange(raw...)
	s.Opposite = true
	return
}

var (
	// SOL constant
	SOL = &Symbol{
		Type: SymbolSOL,
	}
	// EOL constant
	EOL = &Symbol{
		Type: SymbolSOL,
	}

	token = map[rune]*Symbol{
		'\\': NewSymbol('\\'),
		'^':  NewSymbol('^'),
		'$':  NewSymbol('^'),
		'[':  NewSymbol('['),
		']':  NewSymbol(']'),
		'd':  NewSymbolRange([]rune{'0', '9'}),
		'D':  NewSymbolRangeR([]rune{'0', '9'}),
		'f':  NewSymbol('\f'),
		'n':  NewSymbol('\n'),
		'r':  NewSymbol('\r'),
		't':  NewSymbol('\t'),
		'v':  NewSymbol('\v'),
		's':  NewSymbol('\f', '\n', '\r', '\t', 'v'),
		'S':  NewSymbolR('\f', '\n', '\r', '\t', 'v'),
		'w':  NewSymbolRange([]rune{'0', '9'}, []rune{'a', 'z'}, []rune{'A', 'Z'}),
		'W':  NewSymbolRangeR([]rune{'0', '9'}, []rune{'a', 'z'}, []rune{'A', 'Z'}),
	}

	opOr        = NewOp('|')
	opNoOrMore  = NewOp('*')
	opOneOrMore = NewOp('+')
	opNoOrOne   = NewOp('?')
)

// NewOp of regexp
func NewOp(raw rune) (s *Symbol) {
	s = &Symbol{
		Type: SymbolOp,
		Op:   raw,
	}
	return
}
