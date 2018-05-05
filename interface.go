package com

import (
	"fmt"
	"reflect"
	"strconv"
)

//ItoString convert interface to string.
func ItoString(i interface{}) string {
	if i == nil {
		return ""
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.String:
		return i.(string)
	default:
		return fmt.Sprintf("%v", i)
	}
}

//ItoUint64 convert interface to uint64.
func ItoUint64(i interface{}) (uint64, error) {
	if i == nil {
		return 0, fmt.Errorf("arg interface is null")
	}
	k := reflect.TypeOf(i)
	switch k.Kind() {
	case reflect.Uint:
		return uint64(i.(uint)), nil
	case reflect.Uint8:
		return uint64(i.(uint8)), nil
	case reflect.Uint16:
		return uint64(i.(uint16)), nil
	case reflect.Uint32:
		return uint64(i.(uint32)), nil
	case reflect.Uint64:
		return i.(uint64), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := ItoInt64(i)
		return uint64(v), err
	case reflect.Float32, reflect.Float64:
		v, err := ItoFloat64(i)
		return uint64(v), err
	case reflect.String:
		if v, err := strconv.ParseUint(i.(string), 10, 64); err == nil {
			return v, nil
		} else {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind())
	}
}

//ItoUint64Must must convert interface to uint64.
func ItoUint64Must(i interface{}) uint64 {
	v, _ := ItoUint64(i)
	return v
}

//ItoInt convert interface to int.
func ItoInt(i interface{}) (int, error) {
	if i == nil {
		return 0, fmt.Errorf("arg interface value is null")
	}

	k := reflect.TypeOf(i)
	switch k.Kind() {
	case reflect.Int:
		return i.(int), nil
	case reflect.Int8:
		return int(i.(int8)), nil
	case reflect.Int16:
		return int(i.(int16)), nil
	case reflect.Int32:
		return int(i.(int32)), nil
	case reflect.Int64:
		return int(i.(int64)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := ItoUint64(i)
		return int(v), err
	case reflect.Float32, reflect.Float64:
		v, err := ItoFloat64(i)
		return int(v), err
	case reflect.String:
		if v, err := strconv.ParseInt(i.(string), 10, 64); err == nil {
			return int(v), nil
		} else {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind())
	}
}

//ItoIntMust must convert interface to int.
func ItoIntMust(i interface{}) int {
	v, _ := ItoInt(i)
	return v
}

//ItoInt64 convert interface to int64.
func ItoInt64(i interface{}) (int64, error) {
	if i == nil {
		return 0, fmt.Errorf("arg interface value is null")
	}

	k := reflect.TypeOf(i)
	switch k.Kind() {
	case reflect.Int:
		return int64(i.(int)), nil
	case reflect.Int8:
		return int64(i.(int8)), nil
	case reflect.Int16:
		return int64(i.(int16)), nil
	case reflect.Int32:
		return int64(i.(int32)), nil
	case reflect.Int64:
		return i.(int64), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := ItoUint64(i)
		return int64(v), err
	case reflect.Float32, reflect.Float64:
		v, err := ItoFloat64(i)
		return int64(v), err
	case reflect.String:
		if v, err := strconv.ParseInt(i.(string), 10, 64); err == nil {
			return v, nil
		} else {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind())
	}
}

//ItoInt64Must must convert interface to int64.
func ItoInt64Must(i interface{}) int64 {
	v, _ := ItoInt64(i)
	return v
}

//ItoFloat64 convert interface to float64.
func ItoFloat64(i interface{}) (float64, error) {
	if i == nil {
		return 0, fmt.Errorf("arg interface value is null")
	}
	k := reflect.TypeOf(i)
	switch k.Kind() {
	case reflect.Float32:
		return float64(i.(float32)), nil
	case reflect.Float64:
		return float64(i.(float64)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := ItoUint64(i)
		return float64(v), err
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := ItoInt64(i)
		return float64(v), err
	case reflect.String:
		if v, err := strconv.ParseFloat(i.(string), 64); err == nil {
			return v, nil
		} else {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind())
	}
}

//ItoFloat64Must must convert interface to float64.
func ItoFloat64Must(i interface{}) float64 {
	v, _ := ItoFloat64(i)
	return v
}

//ItoMap convert interface to Map.
func ItoMap(i interface{}) (Map, error) {
	if i == nil {
		return nil, nil
	}
	if m, ok := i.(Map); ok {
		return m, nil
	} else if m, ok := i.(map[string]interface{}); ok {
		return Map(m), nil
	} else {
		return nil, fmt.Errorf("incompactable value(%v)", i)
	}
}

//ItoMapMust must convert interface to Map.
func ItoMapMust(i interface{}) Map {
	m, _ := ItoMap(i)
	return m
}

//ItoSlice convert interface to interface slice.
func ItoSlice(i interface{}) []interface{} {
	if vals, ok := i.([]interface{}); ok {
		return vals
	}
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice {
		return nil
	}
	vals := []interface{}{}
	for j := 0; j < v.Len(); j++ {
		vals = append(vals, v.Index(j).Interface())
	}
	return vals
}

//ItoStringSlice convert interface to string slice.
func ItoStringSlice(i interface{}) []string {
	vals := ItoSlice(i)
	if vals == nil {
		return nil
	}
	ss := []string{}
	for _, v := range vals {
		ss = append(ss, ItoString(v))
	}
	return ss
}

//ItoIntSlice convert interface to int slice.
func ItoIntSlice(i interface{}) ([]int, error) {
	vals := ItoSlice(i)
	if vals == nil {
		return nil, nil
	}
	is := []int{}
	for _, v := range vals {
		iv, err := ItoInt(v)
		if err != nil {
			return nil, err
		}
		is = append(is, iv)
	}
	return is, nil
}

//ItoIntSliceMust  must convert interface to int slice.
func ItoIntSliceMust(i interface{}) []int {
	v, _ := ItoIntSlice(i)
	return v
}

//ItoInt64Slice convert interface to int64 slice.
func ItoInt64Slice(i interface{}) ([]int64, error) {
	vals := ItoSlice(i)
	if vals == nil {
		return nil, nil
	}
	is := []int64{}
	for _, v := range vals {
		iv, err := ItoInt64(v)
		if err != nil {
			return nil, err
		}
		is = append(is, iv)
	}
	return is, nil
}

//ItoInt64SliceMust convert interface to int64 slice.
func ItoInt64SliceMust(i interface{}) []int64 {
	v, _ := ItoInt64Slice(i)
	return v
}

//ItoFloat64Slice convert interface to float64 slice.
func ItoFloat64Slice(i interface{}) ([]float64, error) {
	vals := ItoSlice(i)
	if vals == nil {
		return nil, nil
	}
	fs := []float64{}
	for _, v := range vals {
		fv, err := ItoFloat64(v)
		if err != nil {
			return nil, err
		}
		fs = append(fs, fv)
	}
	return fs, nil
}

//ItoFloat64SliceMust convert interface to float64 slice.
func ItoFloat64SliceMust(i interface{}) []float64 {
	v, _ := ItoFloat64Slice(i)
	return v
}

//ItoMapSlice convert interface to Map slice.
func ItoMapSlice(i interface{}) ([]Map, error) {
	vals := ItoSlice(i)
	if vals == nil {
		return nil, nil
	}
	ms := []Map{}
	for _, v := range vals {
		mv, err := ItoMap(v)
		if err != nil {
			return nil, err
		}
		ms = append(ms, mv)
	}
	return ms, nil
}

//ItoMapSliceMust must convert interface to Map slice.
func ItoMapSliceMust(i interface{}) []Map {
	v, _ := ItoMapSlice(i)
	return v
}
