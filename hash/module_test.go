package hash

import (
	"URLShortenerDemo/errors"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestModule(t *testing.T) {
	log := logrus.New()
	hash := New(log)
	Convey("test hash method", t, func() {
		res, err := hash.IDtoShortenID(12)
		So(err, ShouldEqual, nil)
		So(res, ShouldEqual, "lM")

		id, err := hash.ShortenIDtoID("lM")
		So(err, ShouldEqual, nil)
		So(id, ShouldEqual, 12)

		id, err = hash.ShortenIDtoID("abc")
		So(err, ShouldEqual, errors.UrlNotFoundError)
		So(id, ShouldEqual, 0)
	})
}
