package gog

import (
	"errors"
	"testing"
)

func TestIf(t *testing.T) {
	{
		i1, i2 := 1, 2
		exp, got := i1, If(true, i1, i2)
		if got != exp {
			t.Errorf("[int] Expected %d, got: %d", exp, got)
		}
		exp, got = i2, If(false, i1, i2)
		if got != exp {
			t.Errorf("[int] Expected %d, got: %d", exp, got)
		}
	}

	{
		s1, s2 := "first", "second"
		exp, got := s1, If(true, s1, s2)
		if got != exp {
			t.Errorf("[string] Expected %s, got: %s", exp, got)
		}
		exp, got = s2, If(false, s1, s2)
		if got != exp {
			t.Errorf("[string] Expected %s, got: %s", exp, got)
		}
	}
}

func TestCoalesce(t *testing.T) {
	if "" != Coalesce("", "", "") {
		t.Errorf("All args are zero value")
	}
	if "stopHere" != Coalesce("", "stopHere", "") {
		t.Errorf("One arg is not zero value")
	}
	if 123 != Coalesce(0, 0, 123, 432) {
		t.Errorf("More args are not zero value")
	}
	if true != Coalesce(false, false, true) {
		t.Errorf("Bool args are not zero value")
	}
}

func TestPtr(t *testing.T) {
	s := "a"
	sp := Ptr(s)
	if *sp != s {
		t.Errorf("Ptr[string] failed")
	}

	i := 2
	ip := Ptr(i)
	if *ip != i {
		t.Errorf("Ptr[int] failed")
	}
}

func TestMust(t *testing.T) {
	i := 1
	if got := Must(i, nil); got != i {
		t.Errorf("Must[int] failed")
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic")
			}
		}()
		Must(i, errors.New("test")) // Expecting panic
		t.Error("Not expected to reach this")
	}()
}

func manyResults() (i, j, k int, s string, f float64) {
	return 1, 2, 3, "four", 5.0
}

func TestFirst(t *testing.T) {
	exp, got := 1, First(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}

func TestSecond(t *testing.T) {
	exp, got := 2, Second(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}
func TestThird(t *testing.T) {
	exp, got := 3, Third(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}
