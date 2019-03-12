package re

import (
	"io"

	"github.com/ddosakura/gRegEx/core"
	"github.com/ddosakura/gRegEx/fa"
)

// MatchString test str
func (r *Regex) MatchString(str string) bool {
	ri := fa.NewStringInputer(str)
	return core.Match(r.fa, ri)
}

// MatchReader test stream
func (r *Regex) MatchReader(rr io.RuneReader) bool {
	ri := fa.NewReaderInputer(rr)
	return core.Match(r.fa, ri)
}
