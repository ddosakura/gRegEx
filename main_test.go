package re

import (
	"regexp"
	"testing"
)

func testMatchString(t *testing.T, expr string, target string) {
	reg := regexp.MustCompile(expr)
	re := MustCompile(expr)

	a := reg.MatchString(target)
	b := re.MatchString(target)
	t.Log(a, b)

	if a != b {
		t.Errorf("error: /%s/ '%s'", expr, target)
	}
}

func TestA(t *testing.T) {
	testMatchString(t, `a|b`, "ac")
	testMatchString(t, `ab`, "ac")
	testMatchString(t, `a*`, "ac")
}
