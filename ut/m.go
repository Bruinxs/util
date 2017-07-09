package ut

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bruinxs/util/uv"
)

type M map[string]interface{}

func (m M) Exist(key string) bool {
	_, ok := m[key]
	return ok
}

func (m M) Str(key string) string {
	val, ok := m[key]
	if !ok {
		return ""
	}

	sv, err := uv.I2Str(val)
	if err != nil {
		return ""
	}
	return sv
}

func (m M) StrP(patch string) string {
	val := m.ValP(patch)
	if val == nil {
		return ""
	}

	sv, err := uv.I2Str(val)
	if err != nil {
		return ""
	}
	return sv
}

func (m M) Int64(key string) int64 {
	val, ok := m[key]
	if !ok {
		return 0
	}

	iv, err := uv.I2Int64(val)
	if err != nil {
		return 0
	}
	return iv
}

func (m M) Int64P(patch string) int64 {
	val := m.ValP(patch)
	if val == nil {
		return 0
	}

	iv, err := uv.I2Int64(val)
	if err != nil {
		return 0
	}
	return iv
}

func (m M) Int(key string) int {
	return int(m.Int64(key))
}

func (m M) IntP(patch string) int {
	return int(m.Int64P(patch))
}

func (m M) Float(key string) float64 {
	val, ok := m[key]
	if !ok {
		return 0
	}

	fv, err := uv.I2Float64(val)
	if err != nil {
		return 0
	}
	return fv
}

func (m M) FloatP(patch string) float64 {
	val := m.ValP(patch)
	if val == nil {
		return 0
	}

	fv, err := uv.I2Float64(val)
	if err != nil {
		return 0
	}
	return fv
}

func (m M) Bool(key string) bool {
	val, ok := m[key]
	if !ok {
		return false
	}

	bv, err := uv.I2Bool(val)
	if err != nil {
		return false
	}
	return bv
}

func (m M) BoolP(patch string) bool {
	val := m.ValP(patch)
	if val == nil {
		return false
	}

	bv, err := uv.I2Bool(val)
	if err != nil {
		return false
	}
	return bv
}

func (m M) StrSlice(key string) []string {
	val, ok := m[key]
	if !ok {
		return nil
	}

	ssv, err := uv.I2StrSlice(val)
	if err != nil {
		return nil
	}
	return ssv
}

func (m M) StrSliceP(patch string) []string {
	val := m.ValP(patch)
	if val == nil {
		return nil
	}

	ssv, err := uv.I2StrSlice(val)
	if err != nil {
		return nil
	}
	return ssv
}

func (m M) Map(key string) M {
	val, ok := m[key]
	if !ok {
		return nil
	}
	mv, err := uv.I2Map(val)
	if err != nil {
		return nil
	}
	return M(mv)
}

func (m M) MapP(patch string) M {
	val := m.ValP(patch)
	if val == nil {
		return nil
	}
	mv, err := uv.I2Map(val)
	if err != nil {
		return nil
	}
	return M(mv)
}

func (m M) ValP(path string) interface{} {
	path = strings.Trim(path, "/")
	if path == "" {
		return nil
	}

	ps := strings.Split(path, "/")
	val, ok := m[ps[0]]
	if !ok {
		return nil
	}
	vVal := reflect.ValueOf(val)
	for _, key := range ps[1:] {
		kind := vVal.Kind()
		for kind == reflect.Ptr || kind == reflect.Interface {
			if vVal.IsNil() {
				return nil
			}
			vVal = vVal.Elem()
			kind = vVal.Kind()
		}

		switch kind {
		case reflect.Map:
			vVal = vVal.MapIndex(reflect.ValueOf(key))
		case reflect.Slice, reflect.Array:
			vLen := vVal.Len()
			if key == "len" {
				return vLen
			} else if key == "last" {
				if vLen == 0 {
					return nil
				}
				vVal = vVal.Index(vLen - 1)
			} else {
				index, err := strconv.Atoi(key)
				if err != nil {
					panic(fmt.Sprintf("M valp try to convert key(%v) to integer fail with %v by path(%v)", key, err, path))
				}
				vVal = vVal.Index(index)
			}
		case reflect.Struct:
			vVal = vVal.FieldByName(key)
		}
	}

	if !vVal.IsValid() {
		return nil
	}
	return vVal.Interface()
}
