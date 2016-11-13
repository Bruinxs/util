package uv

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Value interface {
	Value(interface{}) interface{}
}

type Fetcher Value

//fetch value by the format string,format define how to get value,it separate rows by ';'
//for a row,separated by ','
//var,m,n
//first option is a variable name
//second, m: must be have the variable, o: it is optional
//third define the range
//'n': no range
//number: a int greater or equal then number,a string length greater or equal then number
//num1~num2: a int greater or equal then num1, and lesser or equal then num2
//val1|val2|val3: the variable value must be val1 or val2 or val3
func Fetch(valuer Value, format string, args ...interface{}) error {
	format = strings.Trim(format, " \r\n\t")
	if format[len(format)-1] == ';' {
		format = format[:len(format)-1]
	}
	rows := strings.Split(format, ";")
	if len(rows) != len(args) {
		return fmt.Errorf("format rows len(%v) not equal arg len(%v)", len(rows), len(args))
	}

	size := len(rows)
	for i := 0; i < size; i++ {
		rows[i] = strings.Trim(rows[i], " \r\n\t")
		options := strings.Split(rows[i], ",")
		if len(options) != 3 {
			return fmt.Errorf("row(%v) option len(%v) not equal 3", rows[i], len(options))
		}

		val := valuer.Value(options[0])
		if val == nil {
			if options[1] == "o" {
				continue
			} else if options[1] == "m" {
				return fmt.Errorf("row(%v) variable(%v) must be provide", rows[i], options[0])
			} else {
				return fmt.Errorf("row(%v) second option(%v) is illegal", rows[i], options[1])
			}
		}

		argVal := reflect.ValueOf(args[i])
		if argVal.Kind() != reflect.Ptr {
			return fmt.Errorf("arg in order %v expect ptr type but a %t", i, args[i])
		}

		elemVal := argVal.Elem()
		switch elemVal.Kind() {
		case reflect.String:
			sv, err := I2Str(val)
			if err != nil {
				return err
			}
			elemVal.SetString(sv)

			err = checkRange(elemVal, options[2])
			if err != nil {
				return fmt.Errorf("row(%v), %v", rows[i], err)
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, err := I2Int64(val)
			if err != nil {
				return err
			}
			elemVal.SetInt(iv)

			err = checkRange(elemVal, options[2])
			if err != nil {
				return fmt.Errorf("row(%v), %v", rows[i], err)
			}

		case reflect.Float32, reflect.Float64:
			fv, err := I2Float64(val)
			if err != nil {
				return err
			}
			elemVal.SetFloat(fv)

			err = checkRange(elemVal, options[2])
			if err != nil {
				return fmt.Errorf("row(%v), %v", rows[i], err)
			}

		case reflect.Bool:
			bv, err := I2Bool(val)
			if err != nil {
				return err
			}
			elemVal.SetBool(bv)

		case reflect.Slice:
			ssv, err := I2StrSlice(val)
			if err != nil {
				return err
			}
			elemVal.Set(reflect.ValueOf(ssv))

		default:
			return fmt.Errorf("row(%v) arg(%v) type(%t) illegal", rows[i], args[i], args[i])
		}
	}
	return nil
}

func checkRange(val reflect.Value, ranges string) error {
	ranges = strings.Trim(ranges, " ")
	if ranges == "" {
		return fmt.Errorf("val(%v) of range is empty", val)
	}
	if ranges == "n" {
		return nil
	}

	var rangeStrV []string
	var rangeFV []float64
	var flg int
	var err error

	if strings.Contains(ranges, "~") {
		rangeStrV = strings.Split(ranges, "~")
		if len(rangeStrV) != 2 {
			return fmt.Errorf("ranges(%v) illegal, should only have two val both ends of '~'", ranges)
		}
		flg = 1
		rangeFV = make([]float64, 2)
		rangeFV[0], err = strconv.ParseFloat(rangeStrV[0], 64)
		if err != nil {
			return fmt.Errorf("ranges(%v) parse str(%v) to float err(%v)", ranges, rangeStrV[0], err)
		}
		rangeFV[1], err = strconv.ParseFloat(rangeStrV[1], 64)
		if err != nil {
			return fmt.Errorf("ranges(%v) parse str(%v) to float err(%v)", ranges, rangeStrV[1], err)
		}
	} else if strings.Contains(ranges, "|") {
		rangeStrV = strings.Split(ranges, "|")
		flg = 2
	} else {
		flg = 3
		rangeFV = make([]float64, 1)
		rangeFV[0], err = strconv.ParseFloat(ranges, 64)
		if err != nil {
			return fmt.Errorf("ranges(%v) parse to float err(%v)", ranges, err)
		}
	}

	switch val.Kind() {
	case reflect.String:
		sv := val.String()
		switch flg {
		case 1, 3:
			if len(sv) < int(rangeFV[0]) {
				return fmt.Errorf("str val(%v) len(%v) less than %v, not in range(%v)", sv, len(sv), rangeFV[0], ranges)
			}
			if flg == 1 && len(sv) > int(rangeFV[1]) {
				return fmt.Errorf("str val(%v) len(%v) greater than %v, not in range(%v)", sv, len(sv), rangeFV[1], ranges)
			}
		case 2:
			ill := true
			for _, s := range rangeStrV {
				if sv == s {
					ill = false
					break
				}
			}
			if ill {
				return fmt.Errorf("str val(%v) not in range(%v)", sv, ranges)
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		iv := val.Int()
		switch flg {
		case 1, 3:
			if iv < int64(rangeFV[0]) {
				return fmt.Errorf("int val(%v) less than %v, not in range(%v)", iv, rangeFV[0], ranges)
			}
			if flg == 1 && iv > int64(rangeFV[1]) {
				return fmt.Errorf("int val(%v) greater than %v, not in range(%v)", iv, rangeFV[1], ranges)
			}
		case 2:
			ill := true
			for _, s := range rangeStrV {
				k, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return fmt.Errorf("parse str(%v) to int err(%v)", s, err)
				}
				if k == iv {
					ill = false
					break
				}
			}
			if ill {
				return fmt.Errorf("int val(%v) not in range(%v)", iv, ranges)
			}
		}
	case reflect.Float32, reflect.Float64:
		fv := val.Float()
		switch flg {
		case 1, 3:
			if fv < rangeFV[0] {
				return fmt.Errorf("float val(%v) less than %v, not in range(%v)", fv, rangeFV[0], ranges)
			}
			if flg == 1 && fv > rangeFV[1] {
				return fmt.Errorf("float val(%v) greater than %v, not in range(%v)", fv, rangeFV[1], ranges)
			}
		case 2:
			ill := true
			for _, s := range rangeStrV {
				f, err := strconv.ParseFloat(s, 64)
				if err != nil {
					return fmt.Errorf("parse str(%v) t float err(%v)", s, err)
				}
				if f == fv {
					ill = false
					break
				}
			}
			if ill {
				return fmt.Errorf("float val(%v) not in range(%v)", fv, ranges)
			}
		}
	}
	return nil
}
