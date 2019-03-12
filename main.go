package re

import (
	"errors"
	"strconv"

	"github.com/ddosakura/gRegEx/core"
	"github.com/ddosakura/gRegEx/fa"
	"github.com/ddosakura/gRegEx/nfa"
)

// FA type
type FA int

const (
	_ FA = iota
	// NFA "github.com/ddosakura/gRegEx/nfa"
	NFA
	// ENFA "github.com/ddosakura/gRegEx/enfa"
	ENFA
	// DFA "github.com/ddosakura/gRegEx/dfa"
	DFA
)

// Regex - Regular Expression
type Regex struct {
	expr string
	fa   fa.FA
}

// Compile str to re
func Compile(expr string, t FA) (regex *Regex, e error) {
	regex = &Regex{
		expr: expr,
		fa:   nil,
	}
	switch t {
	case NFA:
		regex.fa, e = core.NewFA(nfa.New, []rune(expr))
	case ENFA:
		return nil, errors.New("TODO")
	case DFA:
		return nil, errors.New("TODO")
	default:
		return nil, errors.New("ERROR FA TYPE")
	}
	if e != nil {
		regex = nil
	}
	return
}

// MustCompile panic when compile with error
func MustCompile(expr string, t FA) *Regex {
	r, err := Compile(expr, t)
	if err != nil {
		panic(`re: Compile(` + quote(expr) + `): ` + err.Error())
	}
	return r
}

func quote(s string) string {
	if strconv.CanBackquote(s) {
		return "`" + s + "`"
	}
	return strconv.Quote(s)
}
