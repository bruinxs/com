package com

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMap(t *testing.T) {
	Convey("test map get value", t, func() {
		ms := Map{
			"str":   []string{"foo", "bar"},
			"int":   []int{1, 2},
			"float": []float64{1.2, 3.4},
			"map":   []Map{{"foo": "bar"}},
		}

		m := Map{
			"str":   "string",
			"int":   1,
			"float": 1.23,
			"map":   ms,
		}

		Convey("check if exist", func() {
			So(Map(nil).Exist("str"), ShouldBeFalse)
			So(m.Exist("fake"), ShouldBeFalse)
			So(m.Exist("str"), ShouldBeTrue)
		})

		Convey("get string value", func() {
			So(Map(nil).String("str"), ShouldEqual, "")
			So(m.String("fake"), ShouldEqual, "")
			So(m.String("str"), ShouldEqual, "string")
		})

		Convey("get int value", func() {
			So(Map(nil).Int("int"), ShouldEqual, 0)
			So(m.Int("fake"), ShouldEqual, 0)
			So(m.Int("int"), ShouldEqual, 1)
		})

		Convey("get int64 value", func() {
			So(Map(nil).Int64("int"), ShouldEqual, 0)
			So(m.Int64("fake"), ShouldEqual, 0)
			So(m.Int64("int"), ShouldEqual, 1)
		})

		Convey("get float64 value", func() {
			So(Map(nil).Float64("float"), ShouldEqual, 0)
			So(m.Float64("fake"), ShouldEqual, 0)
			So(m.Float64("float"), ShouldEqual, 1.23)
		})

		Convey("get map value", func() {
			So(Map(nil).Map("map"), ShouldBeNil)
			So(m.Map("fake"), ShouldBeNil)
			So(m.Map("map"), ShouldEqual, ms)
		})

		Convey("get string slice value", func() {
			So(Map(nil).StringSlice("str"), ShouldBeNil)
			So(ms.StringSlice("fake"), ShouldBeNil)
			So(ms.StringSlice("str"), ShouldResemble, []string{"foo", "bar"})
		})

		Convey("get int slice value", func() {
			So(Map(nil).IntSlice("int"), ShouldBeNil)
			So(ms.IntSlice("fake"), ShouldBeNil)
			So(ms.IntSlice("int"), ShouldResemble, []int{1, 2})
		})

		Convey("get int64 slice value", func() {
			So(Map(nil).Int64Slice("int"), ShouldBeNil)
			So(ms.Int64Slice("fake"), ShouldBeNil)
			So(ms.Int64Slice("int"), ShouldResemble, []int64{1, 2})
		})

		Convey("get float64 slice value", func() {
			So(Map(nil).Float64Slice("int"), ShouldBeNil)
			So(ms.Float64Slice("fake"), ShouldBeNil)
			So(ms.Float64Slice("float"), ShouldResemble, []float64{1.2, 3.4})
		})

		Convey("get map slice value", func() {
			So(Map(nil).MapSlice("int"), ShouldBeNil)
			So(ms.MapSlice("fake"), ShouldBeNil)
			So(ms.MapSlice("map"), ShouldResemble, []Map{{"foo": "bar"}})
		})
	})
}

func TestMapSet(t *testing.T) {
	Convey("set map value", t, func() {
		m := Map{}
		m.Set("k", 1).Set("k2", "str")
		So(m.Exist("k2"), ShouldBeTrue)
		So(m.Int("k"), ShouldEqual, 1)
		So(m.String("k2"), ShouldEqual, "str")

		m.Set("k2", nil)
		So(m.Exist("k2"), ShouldBeFalse)
	})
}
