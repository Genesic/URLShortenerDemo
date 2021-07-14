package env

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestGetDefault(t *testing.T) {
	Convey("test get default method", t, func() {
		result := GetDefault("test_env", "abc")
		So(result, ShouldEqual, "abc")

		os.Setenv("test_env", "def")
		result = GetDefault("test_env", "abc")
		So(result, ShouldEqual, "def")
	})
}