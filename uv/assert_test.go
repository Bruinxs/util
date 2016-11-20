package uv

import (
	"github.com/Bruinxs/tu"
	"strings"
	"testing"
)

func TestI2Str(t *testing.T) {
	var i interface{}

	i = "string"
	str, err := I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "string" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	i = 10
	str, err = I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "10" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	i = int8(8)
	str, err = I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "8" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	i = int16(16)
	str, err = I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "16" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	i = int32(32)
	str, err = I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "32" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	i = int64(64)
	str, err = I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "64" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	i = 4.56
	str, err = I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "4.56" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	i = float32(3.14)
	str, err = I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "3.14" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	i = true
	str, err = I2Str(i)
	if err != nil {
		t.Error(err)
		return
	}
	if str != "true" {
		t.Errorf("str(%v) != %v", str, i)
		return
	}

	//err
	i = map[string]interface{}{"key": "val"}
	str, err = I2Str(i)
	if err == nil || !strings.Contains(err.Error(), "assert to string fail") {
		t.Error(err)
		return
	}
}

func TestI2Int64(t *testing.T) {
	var i interface{}

	i = 10
	iv, err := I2Int64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if iv != int64(10) {
		t.Errorf("iv(%v) != %v", iv, i)
		return
	}

	i = int8(8)
	iv, err = I2Int64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if iv != int64(8) {
		t.Errorf("iv(%v) != %v", iv, i)
		return
	}

	i = int16(16)
	iv, err = I2Int64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if iv != int64(16) {
		t.Errorf("iv(%v) != %v", iv, i)
		return
	}

	i = int32(32)
	iv, err = I2Int64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if iv != int64(32) {
		t.Errorf("iv(%v) != %v", iv, i)
		return
	}

	i = int64(64)
	iv, err = I2Int64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if iv != int64(64) {
		t.Errorf("iv(%v) != %v", iv, i)
		return
	}

	i = 3.145
	iv, err = I2Int64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if iv != int64(3) {
		t.Errorf("iv(%v) != %v", iv, i)
		return
	}

	i = "1001"
	iv, err = I2Int64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if iv != int64(1001) {
		t.Errorf("iv(%v) != %v", iv, i)
		return
	}

	//err
	i = "fake"
	iv, err = I2Int64(i)
	if err == nil {
		t.Error(err)
		return
	}

	i = true
	iv, err = I2Int64(i)
	if err == nil || !strings.Contains(err.Error(), "assert to int fail") {
		t.Error(err)
		return
	}
}

func TestI2Float64(t *testing.T) {
	var i interface{}

	i = 3.145
	fv, err := I2Float64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if fv != 3.145 {
		t.Errorf("iv(%v) != %v", fv, i)
		return
	}

	i = float32(32.1235)
	fv, err = I2Float64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if fv != float64(float32(32.1235)) {
		t.Errorf("iv(%v) != %v", fv, i)
		return
	}

	i = 64
	fv, err = I2Float64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if fv != 64.0 {
		t.Errorf("iv(%v) != %v", fv, i)
		return
	}

	i = "10.0"
	fv, err = I2Float64(i)
	if err != nil {
		t.Error(err)
		return
	}
	if fv != 10.0 {
		t.Errorf("iv(%v) != %v", fv, i)
		return
	}

	//err
	i = "fake"
	fv, err = I2Float64(i)
	if err == nil {
		t.Error(err)
		return
	}

	i = true
	fv, err = I2Float64(i)
	if err == nil || !strings.Contains(err.Error(), "assert to float fail") {
		t.Error(err)
		return
	}
}

func TestI2Bool(t *testing.T) {
	var i interface{}

	i = true
	bv, err := I2Bool(i)
	if err != nil {
		t.Error(err)
		return
	}
	if bv != true {
		t.Errorf("iv(%v) != %v", bv, i)
		return
	}

	i = 0
	bv, err = I2Bool(i)
	if err != nil {
		t.Error(err)
		return
	}
	if bv != false {
		t.Errorf("iv(%v) != %v", bv, i)
		return
	}

	i = "true"
	bv, err = I2Bool(i)
	if err != nil {
		t.Error(err)
		return
	}
	if bv != true {
		t.Errorf("iv(%v) != %v", bv, i)
		return
	}

	//err
	i = "fake"
	bv, err = I2Bool(i)
	if err == nil {
		t.Error(err)
		return
	}

	i = 10.34
	bv, err = I2Bool(i)
	if err == nil || !strings.Contains(err.Error(), "assert to bool fail") {
		t.Error(err)
		return
	}
}

func TestI2StrSlice(t *testing.T) {
	var i interface{}

	i = []interface{}{"string", 10, 3.14, true}
	ssv, err := I2StrSlice(i)
	if err != nil {
		t.Error(err)
		return
	}
	if !tu.CmpStr_Strict(ssv, []string{"string", "10", "3.14", "true"}) {
		t.Errorf("ssv(%v) != %v", ssv, i)
		return
	}

	i = "s1,s2,s3,s4"
	ssv, err = I2StrSlice(i)
	if err != nil {
		t.Error(err)
		return
	}
	if !tu.CmpStr_Strict(ssv, []string{"s1", "s2", "s3", "s4"}) {
		t.Errorf("ssv(%v) != %v", ssv, i)
		return
	}

	//err
	i = []interface{}{"string", 10, 3.14, true, map[string]string{}}
	_, err = I2StrSlice(i)
	if err == nil {
		t.Error(err)
		return
	}

	i = map[string]string{}
	_, err = I2StrSlice(i)
	if err == nil || !strings.Contains(err.Error(), "assert to string slice fail") {
		t.Error(err)
		return
	}
}
