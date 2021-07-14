package env

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	Convey("test get method", t, func() {
		So(func() { Get("test") }, ShouldPanic)
		os.Setenv("test", "abc")
		result := Get("test")
		So(result, ShouldEqual, "abc")
	})
}

func TestGetDefault(t *testing.T) {
	Convey("test get default method", t, func() {
		result := GetDefault("test_env", "abc")
		So(result, ShouldEqual, "abc")

		os.Setenv("test_env", "def")
		result = GetDefault("test_env", "abc")
		So(result, ShouldEqual, "def")
	})
}

func TestGetDefaultInt(t *testing.T) {
	Convey("test get default int method", t, func() {
		result := GetDefaultInt("test_int", 123)
		So(result, ShouldEqual, 123)

		os.Setenv("test_int", "456")
		result = GetDefaultInt("test_int", 123)
		So(result, ShouldEqual, 456)

		os.Setenv("test_int", "abc")
		So(func() { GetDefaultInt("test_int", 123) }, ShouldPanic)
	})
}