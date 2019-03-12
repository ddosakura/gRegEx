package core

import (
	"github.com/ddosakura/gRegEx/fa"
)

// Parser of regexp
type Parser struct {
	lexer *Lexer
}

// Start parser
func (p *Parser) Start(fn func(s fa.State) fa.FA) (impl fa.FA, e error) {
	startState := fa.NewBasicState(false)
	impl = fn(startState)
	return
}
