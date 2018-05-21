package com

//Map like a object
type Map map[string]interface{}

//Set set value on key, if value is nil then remove it
func (m Map) Set(key string, value interface{}) Map {
	if value == nil {
		delete(m, key)
	} else {
		m[key] = value
	}
	return m
}

//Exist check if exist the key
func (m Map) Exist(key string) bool {
	if m == nil {
		return false
	}
	_, ok := m[key]
	return ok
}

//String get string value by key
func (m Map) String(key string) string {
	if m == nil {
		return ""
	}
	if val, ok := m[key]; ok {
		return ItoString(val)
	}
	return ""
}

//Int get int value by key
func (m Map) Int(key string) int {
	if m == nil {
		return 0
	}
	if val, ok := m[key]; ok {
		return ItoIntMust(val)
	}
	return 0
}

//Int64 get int64 value by key
func (m Map) Int64(key string) int64 {
	if m == nil {
		return 0
	}
	if val, ok := m[key]; ok {
		return ItoInt64Must(val)
	}
	return 0
}

//Float64 get float64 value by key
func (m Map) Float64(key string) float64 {
	if m == nil {
		return 0
	}
	if val, ok := m[key]; ok {
		return ItoFloat64Must(val)
	}
	return 0
}

//Map get map value by key
func (m Map) Map(key string) Map {
	if m == nil {
		return nil
	}
	if val, ok := m[key]; ok {
		return ItoMapMust(val)
	}
	return nil
}

//StringSlice get string slice by key
func (m Map) StringSlice(key string) []string {
	if m == nil {
		return nil
	}
	if val, ok := m[key]; ok {
		return ItoStringSlice(val)
	}
	return nil
}

//IntSlice get int slice by key
func (m Map) IntSlice(key string) []int {
	if m == nil {
		return nil
	}
	if val, ok := m[key]; ok {
		return ItoIntSliceMust(val)
	}
	return nil
}

//Int64Slice get int64 slice by key
func (m Map) Int64Slice(key string) []int64 {
	if m == nil {
		return nil
	}
	if val, ok := m[key]; ok {
		return ItoInt64SliceMust(val)
	}
	return nil
}

//Float64Slice get float64 slice by key
func (m Map) Float64Slice(key string) []float64 {
	if m == nil {
		return nil
	}
	if val, ok := m[key]; ok {
		return ItoFloat64SliceMust(val)
	}
	return nil
}

//MapSlice get Map slice by key
func (m Map) MapSlice(key string) []Map {
	if m == nil {
		return nil
	}
	if val, ok := m[key]; ok {
		return ItoMapSliceMust(val)
	}
	return nil
}
