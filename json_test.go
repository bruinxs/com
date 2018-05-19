package com

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJson(t *testing.T) {
	Convey("must convert to json string", t, func() {
		So(Json(nil), ShouldEqual, "null")
		So(Json(1), ShouldEqual, "1")
		So(Json(1.2), ShouldEqual, "1.2")
		So(Json("str"), ShouldEqual, "\"str\"")
		So(Json(true), ShouldEqual, "true")
		So(Json([]string{"a", "b", "c"}), ShouldEqual, "[\"a\",\"b\",\"c\"]")
		So(Json(Map{"key": "val"}), ShouldEqual, "{\"key\":\"val\"}")
	})
}
