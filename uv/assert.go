package uv

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func I2Val(i interface{}) reflect.Value {
	val := reflect.ValueOf(i)
	kind := val.Kind()
	for kind == reflect.Ptr || kind == reflect.Interface {
		val = val.Elem()
		kind = val.Kind()
	}
	return val
}

func I2Str(i interface{}) (string, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.String:
		return val.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%v", val), nil
	}

	return "", fmt.Errorf("interface val(%v) type(%v) assert to string fail", i, val.Kind())
}

func I2Int64(i interface{}) (int64, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int(), nil
	case reflect.Float32, reflect.Float64:
		return int64(val.Float()), nil
	case reflect.String:
		iv, err := strconv.ParseInt(val.String(), 10, 64)
		if err != nil {
			return 0, err
		}
		return iv, nil
	}

	return 0, fmt.Errorf("interface val(%v) type(%v) assert to int fail", i, val.Kind())
}

func I2Float64(i interface{}) (float64, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int()), nil
	case reflect.String:
		fv, err := strconv.ParseFloat(val.String(), 64)
		if err != nil {
			return 0, err
		}
		return fv, nil
	}

	return 0, fmt.Errorf("interface val(%v) type(%v) assert to float fail", i, val.Kind())
}

func I2Bool(i interface{}) (bool, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Bool:
		return val.Bool(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() != 0, nil
	case reflect.String:
		bv, err := strconv.ParseBool(val.String())
		if err != nil {
			return false, err
		}
		return bv, nil
	}

	return false, fmt.Errorf("interface val(%v) type(%v) assert to bool fail", i, val.Kind())
}

func I2StrSlice(i interface{}) ([]string, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Slice:
		ssv := make([]string, val.Len())
		for i, l := 0, val.Len(); i < l; i++ {
			sv, err := I2Str(val.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			ssv[i] = sv
		}
		return ssv, nil

	case reflect.String:
		return strings.Split(val.String(), ","), nil
	}

	return nil, fmt.Errorf("interface val(%v) type(%v) assert to string slice fail", i, val.Kind())
}

func I2Map(i interface{}) (map[string]interface{}, error) {
	val := I2Val(i)
	switch val.Kind() {
	case reflect.Map:
		m := map[string]interface{}{}
		for _, kv := range val.MapKeys() {
			key, err := I2Str(kv.Interface())
			if err != nil {
				return nil, err
			}
			m[key] = val.MapIndex(kv).Interface()
		}
		return m, nil

	case reflect.Struct:
		m := map[string]interface{}{}
		typ := val.Type()
		fieldNum := typ.NumField()
		var key string
		for i := 0; i < fieldNum; i++ {
			field := typ.Field(i)
			fieldVal := val.Field(i)
			if !fieldVal.CanInterface() {
				continue
			}

			tag := field.Tag.Get("json")
			if tag != "" {
				key = strings.Split(tag, ",")[0]
			} else {
				key = field.Name
			}

			m[key] = fieldVal.Interface()
		}
		return m, nil
	}

	return nil, fmt.Errorf("interface val(%v) type(%v) assert to map fail", i, val.Kind())
}
