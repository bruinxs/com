package com

import (
	"testing"

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
			{1, "1"},
			{[]byte("bytes"), "bytes"},
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
	Convey("convert interface to string slice", t, func() {
		So(ItoStringSlice(nil), ShouldBeNil)
		So(ItoStringSlice([]int{1, 2, 3}), ShouldResemble, []string{"1", "2", "3"})
	})
}

func TestItoIntSlice(t *testing.T) {
	Convey("convert interface to int slice", t, func() {

		Convey("nil argument", func() {
			iv, err := ItoIntSlice(nil)
			So(err, ShouldBeNil)
			So(iv, ShouldBeNil)
		})

		Convey("llegal argument", func() {
			iv, err := ItoIntSlice([]string{"1", "2", "3"})
			So(err, ShouldBeNil)
			So(iv, ShouldResemble, []int{1, 2, 3})

		})

		Convey("must convert", func() {
			iv := ItoIntSliceMust([]string{"1", "2", "3"})
			So(iv, ShouldResemble, []int{1, 2, 3})
		})
	})
}

func TestItoInt64Slice(t *testing.T) {
	Convey("convert interface to int64 slice", t, func() {

		Convey("nil argument", func() {
			iv, err := ItoInt64Slice(nil)
			So(err, ShouldBeNil)
			So(iv, ShouldBeNil)
		})

		Convey("llegal argument", func() {
			iv, err := ItoInt64Slice([]string{"1", "2", "3"})
			So(err, ShouldBeNil)
			So(iv, ShouldResemble, []int64{1, 2, 3})

		})

		Convey("must convert", func() {
			iv := ItoInt64SliceMust([]string{"1", "2", "3"})
			So(iv, ShouldResemble, []int64{1, 2, 3})
		})
	})
}

func TestItoFloat64Slice(t *testing.T) {
	Convey("convert interface to float64 slice", t, func() {

		Convey("nil argument", func() {
			iv, err := ItoFloat64Slice(nil)
			So(err, ShouldBeNil)
			So(iv, ShouldBeNil)
		})

		Convey("llegal argument", func() {
			iv, err := ItoFloat64Slice([]string{"1.1", "2.2", "3.3"})
			So(err, ShouldBeNil)
			So(iv, ShouldResemble, []float64{1.1, 2.2, 3.3})

		})

		Convey("must convert", func() {
			iv := ItoFloat64SliceMust([]string{"1.1", "2.2", "3.3"})
			So(iv, ShouldResemble, []float64{1.1, 2.2, 3.3})
		})
	})
}

func TestItoMap(t *testing.T) {
	Convey("convert interface to Map", t, func() {
		Convey("illegal argument", func() {

			mv, err := ItoMap(1)
			So(err, ShouldNotBeNil)
			So(mv, ShouldBeNil)
		})

		Convey("nil argument", func() {
			mv, err := ItoMap(nil)
			So(err, ShouldBeNil)
			So(mv, ShouldBeNil)

		})

		Convey("legal argument", func() {
			m := map[string]interface{}{"foo": "bar"}

			mv, err := ItoMap(m)
			So(err, ShouldBeNil)
			So(mv, ShouldEqual, m)

			mv, err = ItoMap(mv)
			So(err, ShouldBeNil)
			So(mv, ShouldEqual, m)
		})
	})
}

func TestItoMapSlice(t *testing.T) {
	Convey("convert interface to Map slice", t, func() {
		Convey("nil argument", func() {
			mv, err := ItoMapSlice(nil)
			So(err, ShouldBeNil)
			So(mv, ShouldBeNil)
		})

		Convey("llegal argument", func() {
			mv, err := ItoMapSlice([]map[string]interface{}{nil, {"foo": "bar"}})
			So(err, ShouldBeNil)
			So(mv, ShouldResemble, []Map{nil, {"foo": "bar"}})

		})

		Convey("must convert", func() {
			mv := ItoMapSliceMust([]map[string]interface{}{nil, {"foo": "bar"}})
			So(mv, ShouldResemble, []Map{nil, {"foo": "bar"}})
		})
	})
}
