package core

import (
	"testing"

	"github.com/ddosakura/gRegEx/fa"
)

func TestSymbol(t *testing.T) {
	if !token['t'].Eq(fa.NewRuneInput('\t')) {
		t.Errorf("\\t should equal!\n")
	}
	if token['S'].Eq(fa.NewRuneInput('\t')) {
		t.Errorf("\\t should not equal!\n")
	}
}
