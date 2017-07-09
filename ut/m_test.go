package ut_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bruinxs/ts"
	. "github.com/bruinxs/util/ut"
)

func TestM(t *testing.T) {
	m := M{}

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

	m["s1"] = M{}
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

	m["i1"] = M{}
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

	m["f1"] = M{}
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

	m["b1"] = M{}
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

	m["sl1"] = M{}
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

func TestM_StrP(t *testing.T) {
	type Tmp struct {
		Key1 string
		Key2 string
		Key3 int
		Key4 interface{}
	}

	type args struct {
		patch string
	}
	tests := []struct {
		name string
		m    M
		args args
		want string
	}{
		{"", M{"k1": "v1"}, args{"k1"}, "v1"},
		{"", M{"k1": "v1"}, args{"/k1"}, "v1"},
		{"", M{"k1": M{"k2": "v2"}}, args{"k1/k2"}, "v2"},
		{"", M{"k1": M{"k2": "v2"}}, args{"k1/k2/"}, "v2"},
		{"", M{"k1": &M{"k2": "v2"}}, args{"k1/k2/"}, "v2"},
		{"", M{"k1": [3]string{"v1", "v2", "v3"}}, args{"k1/0"}, "v1"},
		{"", M{"k1": [3]string{"v1", "v2", "v3"}}, args{"k1/1"}, "v2"},
		{"", M{"k1": [3]string{"v1", "v2", "v3"}}, args{"k1/2"}, "v3"},
		{"", M{"k1": [3]string{"v1", "v2", "v3"}}, args{"k1/last"}, "v3"},
		{"", M{"k2": []string{"v1", "v2", "v3"}}, args{"k2/0"}, "v1"},
		{"", M{"k2": []string{"v1", "v2", "v3"}}, args{"k2/1"}, "v2"},
		{"", M{"k2": []string{"v1", "v2", "v3"}}, args{"k2/2"}, "v3"},
		{"", M{"k2": []string{"v1", "v2", "v3"}}, args{"k2/last"}, "v3"},
		{"", M{"k": Tmp{"v1", "v2", 3, nil}}, args{"k/Key1"}, "v1"},
		{"", M{"k": Tmp{"v1", "v2", 3, nil}}, args{"k/Key3"}, "3"},
		{"", M{"k": &Tmp{"v1", "v2", 3, nil}}, args{"k/Key2"}, "v2"},
		{"", M{"k": []M{{"arr": []interface{}{interface{}(&Tmp{"", "", 0, M{"kk": "kkv"}})}}}}, args{"k/last/arr/0/Key4/kk"}, "kkv"},
		{"", M{"k": Tmp{"", "", 0, &Tmp{"", "", 0, M{"k2": "vs"}}}}, args{"k/Key4/Key4/k2"}, "vs"},
		{"", M{"k": Tmp{"", "", 0, &Tmp{"", "", 0, M{}}}}, args{"k/Key4/Key4/k2/k3/k4"}, ""},
	}
	for _, tt := range tests {
		tt.name = fmt.Sprintf("run %v", tt.m)
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.StrP(tt.args.patch); got != tt.want {
				t.Errorf("M.StrP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestM_Int64P(t *testing.T) {
	type args struct {
		patch string
	}
	tests := []struct {
		name string
		m    M
		args args
		want int64
	}{
		{"0", M{"k": 123}, args{"k"}, 123},
		{"1", M{"k": []int64{1, 2, 3}}, args{"k/last"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Int64P(tt.args.patch); got != tt.want {
				t.Errorf("M.Int64P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestM_IntP(t *testing.T) {
	type args struct {
		patch string
	}
	tests := []struct {
		name string
		m    M
		args args
		want int
	}{
		{"0", M{"k": 1}, args{"k"}, 1},
		{"1", M{"k": []int{0, 1, 2}}, args{"k/1"}, 1},
		{"2", M{"k": []int{0, 1, 2}}, args{"k/len"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IntP(tt.args.patch); got != tt.want {
				t.Errorf("M.IntP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestM_FloatP(t *testing.T) {
	type args struct {
		patch string
	}
	tests := []struct {
		name string
		m    M
		args args
		want float64
	}{
		{"0", M{"k": 1.23}, args{"k"}, 1.23},
		{"1", M{"k": []float64{1.1, 2.2}}, args{"k/0"}, 1.1},
		{"2", M{"k": []float64{1.1, 2.2}}, args{"k/last"}, 2.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.FloatP(tt.args.patch); got != tt.want {
				t.Errorf("M.FloatP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestM_BoolP(t *testing.T) {
	type args struct {
		patch string
	}
	tests := []struct {
		name string
		m    M
		args args
		want bool
	}{
		{"0", M{"k": true}, args{"k"}, true},
		{"1", M{"k": struct{ I int }{1}}, args{"k/I"}, true},
		{"2", M{"k": []string{"true"}}, args{"k/0"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.BoolP(tt.args.patch); got != tt.want {
				t.Errorf("M.BoolP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestM_StrSliceP(t *testing.T) {
	type args struct {
		patch string
	}
	tests := []struct {
		name string
		m    M
		args args
		want []string
	}{
		{"0", M{"k": "v1"}, args{"k"}, []string{"v1"}},
		{"1", M{"k": "v1,v2,v3"}, args{"k"}, []string{"v1", "v2", "v3"}},
		{"2", M{"k": []string{"v1", "v2"}}, args{"k"}, []string{"v1", "v2"}},
		{"3", M{"k": [][]string{[]string{"v4", "v5"}}}, args{"k/0"}, []string{"v4", "v5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.StrSliceP(tt.args.patch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("M.StrSliceP() = %v, want %v", got, tt.want)
			}
		})
	}
}
