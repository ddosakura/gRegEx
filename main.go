package re

import "strconv"

// Regex - Regular Expression
type Regex struct {
	expr    string
	enginer *engine
}

// Compile str to re
func Compile(expr string) (*Regex, error) {
	enginer, e := newEngineNFA(expr)
	if e != nil {
		return nil, e
	}
	return &Regex{
		expr,
		enginer,
	}, nil
}

// MustCompile panic when compile with error
func MustCompile(expr string) *Regex {
	r, err := Compile(expr)
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
