package ut

import "github.com/bruinxs/util/uv"

type M map[string]interface{}

func (m M) Exist(key string) bool {
	_, ok := m[key]
	return ok
}

func (m M) StrV(key string) string {
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

func (m M) Int64V(key string) int64 {
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

func (m M) IntV(key string) int {
	return int(m.Int64V(key))
}

func (m M) FloatV(key string) float64 {
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

func (m M) BoolV(key string) bool {
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

func (m M) StrSliceV(key string) []string {
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
