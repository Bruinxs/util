package ut_test

import (
	"testing"

	"github.com/Bruinxs/util/ut"
	"github.com/bruinxs/ts"
)

func TestM(t *testing.T) {
	m := ut.M{}

	//exist
	if g, w := m.Exist("fake"), false; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["real"] = 1
	if g, w := m.Exist("real"), true; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	//string
	if g, w := m.Str("s1"), ""; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["s1"] = ut.M{}
	if g, w := m.Str("s1"), ""; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["s1"] = "string"
	if g, w := m.Str("s1"), "string"; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	//int
	if g, w := m.Int("i1"), 0; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["i1"] = ut.M{}
	if g, w := m.Int("i1"), 0; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["i1"] = 10
	if g, w := m.Int("i1"), 10; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	//float
	if g, w := m.Float("f1"), 0.0; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["f1"] = ut.M{}
	if g, w := m.Float("f1"), 0.0; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["f1"] = 3.1415
	if g, w := m.Float("f1"), 3.1415; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	//bool
	if g, w := m.Bool("b1"), false; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["b1"] = ut.M{}
	if g, w := m.Bool("b1"), false; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["b1"] = true
	if g, w := m.Bool("b1"), true; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["b1"] = "false"
	if g, w := m.Bool("b1"), false; g != w {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	//string slice
	if g := m.StrSlice("sl1"); g != nil {
		t.Errorf("got(%v) != %v", g, nil)
		return
	}

	m["sl1"] = ut.M{}
	if g := m.StrSlice("sl1"); g != nil {
		t.Errorf("got(%v) != %v", g, nil)
		return
	}

	m["sl1"] = []string{"s1", "s2", "s3"}
	if g, w := m.StrSlice("sl1"), []string{"s1", "s2", "s3"}; !ts.CmpStr_Strict(g, w) {
		t.Errorf("got(%v) != %v", g, w)
		return
	}

	m["sl1"] = "s4,s5,s6"
	if g, w := m.StrSlice("sl1"), []string{"s4", "s5", "s6"}; !ts.CmpStr_Strict(g, w) {
		t.Errorf("got(%v) != %v", g, w)
		return
	}
}
