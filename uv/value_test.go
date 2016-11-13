package uv

import (
	"fmt"
	"github.com/Bruinxs/tu/ts"
	"reflect"
	"strings"
	"testing"
)

type MV map[string]interface{}

func (mv MV) Value(key interface{}) interface{} {
	if str, ok := key.(string); ok {
		return mv[str]
	}
	panic(fmt.Sprintf("key(%T) ill", key))
}

func TestCheckRange(t *testing.T) {
	tvl := []struct {
		Val   interface{}
		Range string
		Ok    bool
	}{
		{nil, "  ", false},
		{nil, "n", true},
		{nil, "1~2~3", false},
		{nil, "s1~2", false},
		{nil, "1~s2", false},
		{nil, "str", false},
		//
		{10, "10", true},
		{10, "11", false},
		{10, "10~11", true},
		{10, "9~10", true},
		{10, "1~9", false},
		{10, "11~19", false},
		{10, "1|2|10", true},
		{10, "1|2|100", false},
		{10, "s1|s2|s3", false},
		{10.98, "10.98", true},
		{10.98, "10.99", false},
		{10.98, "10.98~11", true},
		{10.98, "10~10.98", true},
		{10.98, "10~10.97", false},
		{10.98, "11.78~12", false},
		{10.98, "100.87|120.23|10.98", true},
		{10.98, "10.97|10.99|100.01", false},
		{10.98, "s1|s2", false},
		{"str", "3", true},
		{"str", "4", false},
		{"str", "3~4", true},
		{"str", "2~3", true},
		{"str", "0~1", false},
		{"str", "4~90", false},
		{"str fake", "str|str fake", true},
		{"str", "s1|s2|s3", false},
	}
	for _, tv := range tvl {
		err := checkRange(reflect.ValueOf(tv.Val), tv.Range)
		if (err == nil) != tv.Ok {
			t.Errorf("tv(%v), err(%v)", tv, err)
			return
		}
	}
}

func TestFetch(t *testing.T) {
	m := MV{
		"s1": "str",
		"s2": "str2",
		"s4": "str1,str2,str3,str4",
		//
		"i1": 2,
		"f1": 3.14159,
		"b1": true,
		//
		"i8":  int8(8),
		"i16": int16(16),
		"i32": int32(32),
		"i64": int64(64),
		"is":  "100",
		//
		"f32": float32(32.14),
		"fs":  "3.1415",
		//
		"bs": "true",
		//
		"m2": map[string]interface{}{},
	}
	var (
		s1 string
		s3 string = "s3"
		s4 []string
	)

	//1,must
	err := Fetch(m, `
		s3,m,0;
	`, &s3)
	if err == nil || !strings.Contains(err.Error(), "must be provide") {
		t.Errorf("err(%v) is illegal", err)
		return
	}

	//2,range
	err = Fetch(m, "s1,m,4", &s1)
	if err == nil || !strings.Contains(err.Error(), "not in range") {
		t.Errorf("err(%v) is illegal", err)
		return
	}

	//3,ptr
	err = Fetch(m, "s1,m,2", s1)
	if err == nil || !strings.Contains(err.Error(), "expect ptr type") {
		t.Errorf("err(%v) is illegal", err)
		return
	}

	//4,arg
	err = Fetch(m, "s1,m,2")
	if err == nil || !strings.Contains(err.Error(), "not equal arg len") {
		t.Errorf("err(%v) is illegal", err)
		return
	}

	err = Fetch(m, "m2,m,n", &s1)
	if err == nil || !strings.Contains(err.Error(), "assert to string fail") {
		t.Error(err)
		return
	}

	//5
	err = Fetch(m, `
		s1,m,3;
		s3,o,0;
		s4,m,n;
	`, &s1, &s3, &s4)
	if err != nil {
		t.Error(err)
		return
	}
	if got, want := s1, m["s1"].(string); got != want {
		t.Errorf("s1(%v) != %v", got, want)
		return
	}
	if got, want := s3, "s3"; got != want {
		t.Errorf("s3(%v) != %v", got, want)
		return
	}
	if got, want := s4, []string{"str1", "str2", "str3", "str4"}; !ts.CmpStr_Strict(got, want) {
		t.Errorf("s4(%v) != %v", got, want)
		return
	}

	//6
	var (
		i1 int
		f1 float64
		b1 bool
	)

	err = Fetch(m, `
		i1,o,2~3;
		f1,m,2.14|3.14159;
		b1,m,n;
	`, &i1, &f1, &b1)
	if err != nil {
		t.Error(err)
		return
	}

	if got, want := i1, 2; got != want {
		t.Errorf("i1(%v) != %v", got, want)
		return
	}
	if got, want := f1, 3.14159; got != want {
		t.Errorf("f1(%v) != %v", got, want)
		return
	}
	if got, want := b1, true; got != want {
		t.Errorf("b1(%v) != %v", got, want)
		return
	}

	//7
	//string
	err = Fetch(m, "i1,m,1;", &s1)
	if err != nil {
		t.Error(err)
	}
	if got, want := s1, "2"; got != want {
		t.Errorf("s1(%v) != %v", got, want)
		return
	}

	//int
	err = Fetch(m, "i8,m,0;", &i1)
	if err != nil {
		t.Error(err)
	}
	if got, want := i1, 8; got != want {
		t.Errorf("i1(%v) != %v", got, want)
		return
	}

	err = Fetch(m, "i16,m,0;", &i1)
	if err != nil {
		t.Error(err)
	}
	if got, want := i1, 16; got != want {
		t.Errorf("i1(%v) != %v", got, want)
		return
	}

	err = Fetch(m, "i32,m,0;", &i1)
	if err != nil {
		t.Error(err)
	}
	if got, want := i1, 32; got != want {
		t.Errorf("i1(%v) != %v", got, want)
		return
	}

	err = Fetch(m, "i64,m,0;", &i1)
	if err != nil {
		t.Error(err)
	}
	if got, want := i1, 64; got != want {
		t.Errorf("i1(%v) != %v", got, want)
		return
	}

	err = Fetch(m, "is,m,0;", &i1)
	if err != nil {
		t.Error(err)
	}
	if got, want := i1, 100; got != want {
		t.Errorf("i1(%v) != %v", got, want)
		return
	}

	//float
	err = Fetch(m, "f32,m,0;", &f1)
	if err != nil {
		t.Error(err)
	}
	if got, want := f1, float64(float32(32.14)); got != want {
		t.Errorf("f1(%v) != %v", got, want)
		return
	}

	err = Fetch(m, "fs,m,0;", &f1)
	if err != nil {
		t.Error(err)
	}
	if got, want := f1, 3.1415; got != want {
		t.Errorf("f1(%v) != %v", got, want)
		return
	}

	//bool
	b1 = false
	err = Fetch(m, "i1,m,0;", &b1)
	if err != nil {
		t.Error(err)
	}
	if got, want := b1, true; got != want {
		t.Errorf("b1(%v) != %v", got, want)
		return
	}

	b1 = false
	err = Fetch(m, "bs,m,0;", &b1)
	if err != nil {
		t.Error(err)
	}
	if got, want := b1, true; got != want {
		t.Errorf("b1(%v) != %v", got, want)
		return
	}

	//8 option err
	err = Fetch(m, "s1,m;", &s1)
	if err == nil || !strings.Contains(err.Error(), "not equal 3") {
		t.Error(err)
		return
	}

	err = Fetch(m, "t1,f,n;", &s1)
	if err == nil || !strings.Contains(err.Error(), "second option") {
		t.Error(err)
		return
	}

	err = Fetch(m, "s1,m,0;", &i1)
	if err == nil || !strings.Contains(err.Error(), "invalid syntax") {
		t.Error(err)
		return
	}

	err = Fetch(m, "b1,m,0;", &i1)
	if err == nil || !strings.Contains(err.Error(), "assert to int fail") {
		t.Error(err)
		return
	}

	err = Fetch(m, "i1,m,0|1;", &i1)
	if err == nil || !strings.Contains(err.Error(), "row(i1,m,0|1)") {
		t.Error(err)
		return
	}

	err = Fetch(m, "s1,m,n;", &f1)
	if err == nil || !strings.Contains(err.Error(), "invalid syntax") {
		t.Error(err)
		return
	}

	err = Fetch(m, "b1,m,n;", &f1)
	if err == nil || !strings.Contains(err.Error(), "assert to float fail") {
		t.Error(err)
		return
	}

	err = Fetch(m, "f1,m,0|0;", &f1)
	if err == nil || !strings.Contains(err.Error(), "row(f1,m,0|0)") {
		t.Error(err)
		return
	}

	err = Fetch(m, "s1,m,0|0;", &b1)
	if err == nil || !strings.Contains(err.Error(), "invalid syntax") {
		t.Error(err)
		return
	}

	err = Fetch(m, "f1,m,0|0;", &b1)
	if err == nil || !strings.Contains(err.Error(), "assert to bool fail") {
		t.Error(err)
		return
	}

	err = Fetch(m, "f1,m,0|0;", &s4)
	if err == nil || !strings.Contains(err.Error(), "assert to string slice fail") {
		t.Error(err)
		return
	}

	var (
		m1 map[string]int
	)
	err = Fetch(m, "f1,m,0|0;", &m1)
	if err == nil || !strings.Contains(err.Error(), "illegal") {
		t.Error(err)
		return
	}
}
