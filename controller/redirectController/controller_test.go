package redirectController

import (
	"URLShortenerDemo/pkg/errors"
	"URLShortenerDemo/pkg/path"
	"URLShortenerDemo/pkg/testUtils"
	"URLShortenerDemo/service/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

func TestModule_Controller(t *testing.T) {
	Convey("test redirect api", t, func() {
		Convey("should return 404 if miss path parameter", func() {
			ctrl := gomock.NewController(t)
			_, r, _, _ := initModule(ctrl)
			res := testUtils.FireRequest(r, http.MethodGet, "/", nil)
			So(res.Code, ShouldEqual, http.StatusNotFound)
		})

		Convey("should return error if shortenId decode failed", func() {
			ctrl := gomock.NewController(t)
			_, r, hash, _ := initModule(ctrl)
			hash.EXPECT().ShortenIDtoID("abc").Return(uint(0), errors.UrlNotFoundError)
			res := testUtils.FireRequest(r, http.MethodGet, "/abc", nil)
			testUtils.VerifyErrorRes(res, errors.UrlNotFoundError)
		})

		Convey("should return error if fetch url failed", func() {
			ctrl := gomock.NewController(t)
			m, r, hash, url := initModule(ctrl)
			urlObject := genData()
			hash.EXPECT().ShortenIDtoID("abc").Return(urlObject.ID, nil)
			url.EXPECT().Get(m.db, urlObject.ID).Return(nil, errors.FetchDatabaseFailedError)
			res := testUtils.FireRequest(r, http.MethodGet, "/abc", nil)
			testUtils.VerifyErrorRes(res, errors.FetchDatabaseFailedError)
		})

		Convey("should return UrlNotFoundError if url is expired", func() {
			ctrl := gomock.NewController(t)
			m, r, hash, _ := initModule(ctrl)
			url := genData()
			url.ExpiredAt = time.Now().Add(-1 * time.Hour)
			m.cacheClient.Set("abc", url, time.Hour)
			hash.EXPECT().ShortenIDtoID("abc").Return(url.ID, nil)
			res := testUtils.FireRequest(r, http.MethodGet, "/abc", nil)
			testUtils.VerifyErrorRes(res, errors.UrlNotFoundError)
		})


		Convey("should return 302 if catch is hit", func() {
			ctrl := gomock.NewController(t)
			m, r, hash, _ := initModule(ctrl)
			url := genData()
			url.ExpiredAt = url.ExpiredAt.Add(time.Hour)
			m.cacheClient.Set("abc", url, time.Hour)
			hash.EXPECT().ShortenIDtoID("abc").Return(url.ID, nil)
			res := testUtils.FireRequest(r, http.MethodGet, "/abc", nil)
			So(res.Code, ShouldEqual, http.StatusFound)
			So(res.Header().Get("Location"), ShouldEqual, url.Url)
		})

		Convey("should return 302 if catch is not hit but url existed", func() {
			ctrl := gomock.NewController(t)
			m, r, hash, url := initModule(ctrl)
			urlObject := genData()
			hash.EXPECT().ShortenIDtoID("abc").Return(urlObject.ID, nil)
			url.EXPECT().Get(m.db, urlObject.ID).Return(urlObject, nil)
			res := testUtils.FireRequest(r, http.MethodGet, "/abc", nil)
			So(res.Code, ShouldEqual, http.StatusFound)
			So(res.Header().Get("Location"), ShouldEqual, urlObject.Url)

			// test cache set
			row, ok := m.cacheClient.Get("abc")
			So(ok, ShouldEqual, true)
			val, ok := row.(*database.Url)
			So(ok, ShouldEqual, true)
			So(val, ShouldEqual, urlObject)
		})
	})
}

func initModule(ctrl *gomock.Controller) (*Module, *gin.Engine, *MockIHash, *MockIUrl) {
	log := logrus.New()
	db, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{SkipDefaultTransaction: true})
	client := cache.New(1*time.Hour, 2*time.Hour)
	hash := NewMockIHash(ctrl)
	url := NewMockIUrl(ctrl)
	module := New(log, gormDB, client, hash, url)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET(path.RedirectApiPath, module.Controller)

	return module, r, hash, url
}

func genData() *database.Url {
	url := &database.Url{
		ID:        123,
		Url:       "https://example.com",
		ExpiredAt: time.Now().Add(time.Minute),
	}

	return url
}
