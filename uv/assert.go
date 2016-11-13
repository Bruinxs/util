package uv

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

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
