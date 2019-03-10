package re

import (
	"io"
)

// MatchString test str
func (r *Regex) MatchString(str string) bool {
	return r.match(stringInput(str))
}

// MatchReader test stream
func (r *Regex) MatchReader(rr io.RuneReader) bool {
	return r.match(readerInput(rr))
}

type input struct {
	rs  []rune
	pos int
	r   io.RuneReader
}

func stringInput(str string) *input {
	return &input{
		rs:  []rune(str),
		pos: 0,
		r:   nil,
	}
}

func readerInput(r io.RuneReader) *input {
	return &input{
		r: r,
	}
}

func (i *input) nextSymbol() (r rune, err error) {
	if i.r == nil {
		if i.pos == len(i.rs) {
			return 0, io.EOF
		}
		r = i.rs[i.pos]
		i.pos++
		return
	}
	r, _, e := i.r.ReadRune()
	if e != nil {
		return 0, e
	}
	return
}

func (r *Regex) match(i *input) bool {
	r.enginer.init()
	for {
		c, e := i.nextSymbol()
		if e == io.EOF {
			return true
		}
		if e != nil {
			panic(e)
		}

		if r.enginer.next(c) != nil {
			return false
		}
	}
}
