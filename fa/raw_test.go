package fa

import (
	"testing"
)

func TestStringInputer(t *testing.T) {
	si := NewStringInputer("abcdefg")
	for i := 0; i < 7; i++ {
		c := si.Next()
		if c == nil {
			t.Error("should not be nil")
		}
		t.Log(c)
	}
	c := si.Next()
	if c != nil {
		t.Error("should be nil")
	}
	t.Log("EOF")
}
