package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_PASS_GeneralForm(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	s1 := "Value Got/Wnt"
	s2 := "Value Got/Wnt"

	chk.Str(s1, s2)
	chk.Str(s1, s2, "unformatted", " message", " not", " displayed")
	chk.Strf(s1, s2, "formatted %s %s %s", "message", "not", "displayed")
}

func Test_FAIL_GeneralForm(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.FailFast(false) // Do not stop on first problem.

	s1 := "Value Got"
	s2 := "Value Wnt"

	chk.Str(s1, s2)
	chk.Str(s1, s2, "unformatted", " message", " displayed")
	chk.Strf(s1, s2, "formatted %s %s", "message", "displayed")
}
