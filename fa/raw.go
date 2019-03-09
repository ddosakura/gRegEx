package fa

import "io"

// RuneInput of rune
type RuneInput struct {
	r rune
}

// NewRuneInput for rune
func NewRuneInput(r rune) *RuneInput {
	return &RuneInput{
		r,
	}
}

// Eq for RuneInput
func (ri *RuneInput) Eq(i Input) (b bool) {
	defer (func() {
		e := recover()
		if e != nil {
			b = false
		}
	})()
	r := i.(*RuneInput)
	return r.r == ri.r
}

// StringInputer of String
type StringInputer struct {
	runes []rune
	pos   int
}

// NewStringInputer create rune stream
func NewStringInputer(str string) *StringInputer {
	return &StringInputer{
		runes: []rune(str),
		pos:   0,
	}
}

// Next for StringInputer
func (si *StringInputer) Next() (i Input) {
	if si.pos == len(si.runes) {
		return nil
	}
	i = &RuneInput{
		r: si.runes[si.pos],
	}
	si.pos++
	return
}

// ReaderInputer of io.RuneReader
type ReaderInputer struct {
	r io.RuneReader
}

// NewReaderInputer create rune stream
func NewReaderInputer(r io.RuneReader) *ReaderInputer {
	return &ReaderInputer{
		r,
	}
}

// Next for ReaderInputer
func (ri *ReaderInputer) Next() (i Input) {
	r, _, e := ri.r.ReadRune()
	if e == io.EOF {
		return nil
	}
	if e != nil {
		panic(e)
	}

	return &RuneInput{
		r,
	}
}

// BasicState implements State
type BasicState struct {
	final bool
}

// NewBasicState create State
func NewBasicState(final bool) *BasicState {
	return &BasicState{
		final: final,
	}
}

// IsFinal of state
func (bs *BasicState) IsFinal() bool {
	return bs.final
}
