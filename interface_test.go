package com

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/smartystreets/goconvey/convey"
)

func TestItoString(t *testing.T) {
	Convey("interface to string", t, func() {
		for _, arg := range []struct {
			i    interface{}
			want string
		}{
			{"string", "string"},
			{12.34, "12.34"},
			{nil, ""},
		} {
			So(ItoString(arg.i), ShouldEqual, arg.want)
		}
	})

}

func TestItoUinit64(t *testing.T) {
	Convey("convert interface to uint64 with nil error", t, func() {
		for _, arg := range []struct {
			i    interface{}
			want uint64
		}{
			{uint(12), 12},
			{uint8(12), 12},
			{uint16(12), 12},
			{uint32(12), 12},
			{uint64(12), 12},
			{12, 12},
			{int64(12), 12},
			{float32(12.34), 12},
			{12.34, 12},
			{"12", 12},
		} {
			get, err := ItoUint64(arg.i)
			So(err, ShouldBeNil)
			So(get, ShouldEqual, arg.want)
		}
	})

	Convey("convert interface to uint64 with error", t, func() {
		for _, i := range []interface{}{
			nil, "xx", So,
		} {
			get, err := ItoUint64(i)
			So(err, ShouldNotBeNil)
			So(get, ShouldEqual, 0)
		}
	})

	Convey("must convert interface to uint64", t, func() {
		for _, arg := range []struct {
			i    interface{}
			want uint64
		}{
			{uint(12), 12},
			{uint8(12), 12},
			{uint16(12), 12},
			{uint32(12), 12},
			{uint64(12), 12},
			{12, 12},
			{int64(12), 12},
			{float32(12.34), 12},
			{12.34, 12},
			{"12", 12},
			{nil, 0},
			{"xx", 0},
			{So, 0},
		} {
			get := ItoUint64Must(arg.i)
			So(get, ShouldEqual, arg.want)
		}
	})
}

func TestItoInt(t *testing.T) {
	Convey("convert interface to int with nil error", t, func() {
		for _, arg := range []struct {
			i    interface{}
			want int
		}{
			{uint(12), 12},
			{uint8(12), 12},
			{uint16(12), 12},
			{uint32(12), 12},
			{uint64(12), 12},
			{12, 12},
			{int64(12), 12},
			{float32(12.34), 12},
			{12.34, 12},
			{"12", 12},
		} {
			get, err := ItoInt(arg.i)
			So(err, ShouldBeNil)
			So(get, ShouldEqual, arg.want)
		}
	})

	Convey("convert interface to int with error", t, func() {
		for _, i := range []interface{}{
			nil, "xx", So,
		} {
			get, err := ItoInt(i)
			So(err, ShouldNotBeNil)
			So(get, ShouldEqual, 0)
		}
	})

	Convey("must convert interface to int", t, func() {
		for _, arg := range []struct {
			i    interface{}
			want int
		}{
			{uint(12), 12},
			{uint8(12), 12},
			{uint16(12), 12},
			{uint32(12), 12},
			{uint64(12), 12},
			{12, 12},
			{int64(12), 12},
			{float32(12.34), 12},
			{12.34, 12},
			{"12", 12},
			{nil, 0},
			{"xx", 0},
			{So, 0},
		} {
			get := ItoIntMust(arg.i)
			So(get, ShouldEqual, arg.want)
		}
	})
}

func TestItoFloat64(t *testing.T) {
	Convey("convert interface to float64 with nil error", t, func() {
		for _, arg := range []struct {
			i    interface{}
			want float64
		}{
			{uint(12), 12},
			{uint8(12), 12},
			{uint16(12), 12},
			{uint32(12), 12},
			{uint64(12), 12},
			{12, 12},
			{int64(12), 12},
			{"12.34", 12.34},
		} {
			get, err := ItoFloat64(arg.i)
			So(err, ShouldBeNil)
			So(get, ShouldEqual, arg.want)
		}
	})

	Convey("convert interface to float64 with error", t, func() {
		for _, i := range []interface{}{
			nil, "xx", So,
		} {
			get, err := ItoFloat64(i)
			So(err, ShouldNotBeNil)
			So(get, ShouldEqual, 0)
		}
	})

	Convey("must convert interface to float64", t, func() {
		for _, arg := range []struct {
			i    interface{}
			want float64
		}{
			{uint(12), 12},
			{uint8(12), 12},
			{uint16(12), 12},
			{uint32(12), 12},
			{uint64(12), 12},
			{12, 12},
			{int64(12), 12},
			{"12.34", 12.34},
			{nil, 0},
			{"xx", 0},
			{So, 0},
		} {
			get := ItoFloat64Must(arg.i)
			So(get, ShouldEqual, arg.want)
		}
	})
}

func TestItoStringSlice(t *testing.T) {
	var ss []string
	assert.Equal(t, ItoStringSlice(nil), ss)
	assert.ElementsMatch(t, ItoStringSlice([]int{1, 2, 3}), []string{"1", "2", "3"})
}

func TestItoIntSlice(t *testing.T) {
	var ii []int
	is, err := ItoIntSlice(nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, is, ii)

	is, err = ItoIntSlice([]string{"1", "2", "3"})
	assert.Equal(t, err, nil)
	assert.ElementsMatch(t, is, []int{1, 2, 3})

	is = ItoIntSliceMust([]string{"1", "2", "3"})
	assert.ElementsMatch(t, is, []int{1, 2, 3})
}

func TestItoInt64Slice(t *testing.T) {
	var ii []int64
	is, err := ItoInt64Slice(nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, is, ii)

	is, err = ItoInt64Slice([]string{"1", "2", "3"})
	assert.Equal(t, err, nil)
	assert.ElementsMatch(t, is, []int64{1, 2, 3})

	is = ItoInt64SliceMust([]string{"1", "2", "3"})
	assert.ElementsMatch(t, is, []int64{1, 2, 3})
}

func TestItoFloat64Slice(t *testing.T) {
	var fi []float64
	fs, err := ItoFloat64Slice(nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, fs, fi)

	fs, err = ItoFloat64Slice([]string{"1.1", "2.2", "3.3"})
	assert.Equal(t, err, nil)
	assert.ElementsMatch(t, fs, []float64{1.1, 2.2, 3.3})

	fs = ItoFloat64SliceMust([]string{"1.1", "2.2", "3.3"})
	assert.ElementsMatch(t, fs, []float64{1.1, 2.2, 3.3})
}
