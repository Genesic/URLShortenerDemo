package shortenController

import (
	"URLShortenerDemo/pkg/errors"
	"URLShortenerDemo/pkg/path"
	"URLShortenerDemo/pkg/testUtils"
	"URLShortenerDemo/service/database"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

func TestModule_Controller(t *testing.T) {
	Convey("test shorten api", t, func() {
		Convey("should return ValidateRequestFailedError if bad request", func() {
			ctrl := gomock.NewController(t)
			_, r, _, _ := initModule(ctrl)
			req, _, _ := genData()
			req.Url = ""
			res := testUtils.FireRequest(r, http.MethodPost, path.ShortenApiPath, req)
			testUtils.VerifyErrorRes(res, errors.ValidateRequestFailedError)

			req, _, _ = genData()
			req.Url = "%"
			res = testUtils.FireRequest(r, http.MethodPost, path.ShortenApiPath, req)
			testUtils.VerifyErrorRes(res, errors.ValidateRequestFailedError)

			req, _, _ = genData()
			req.ExpiredAt = ""
			res = testUtils.FireRequest(r, http.MethodPost, path.ShortenApiPath, req)
			testUtils.VerifyErrorRes(res, errors.ValidateRequestFailedError)

			req, _, _ = genData()
			req.ExpiredAt = "2021-02-08T09:20:041Z"
			res = testUtils.FireRequest(r, http.MethodPost, path.ShortenApiPath, req)
			testUtils.VerifyErrorRes(res, errors.ValidateRequestFailedError)
		})

		Convey("should return error if create url failed", func() {
			ctrl := gomock.NewController(t)
			m, r, url, _ := initModule(ctrl)
			req, expiredAt, _ := genData()
			url.EXPECT().Create(m.db, req.Url, expiredAt).Return(nil, errors.InsertDataBaseFailedError)

			res := testUtils.FireRequest(r, http.MethodPost, path.ShortenApiPath, req)
			testUtils.VerifyErrorRes(res, errors.InsertDataBaseFailedError)
		})

		Convey("should return error if hash shortenId failed", func() {
			ctrl := gomock.NewController(t)
			m, r, url, hash := initModule(ctrl)
			req, expiredAt, urlObject := genData()
			url.EXPECT().Create(m.db, req.Url, expiredAt).Return(urlObject, nil)
			hash.EXPECT().IDtoShortenID(urlObject.ID).Return("", errors.EncodeHashFailedError)

			res := testUtils.FireRequest(r, http.MethodPost, path.ShortenApiPath, req)
			testUtils.VerifyErrorRes(res, errors.EncodeHashFailedError)
		})

		Convey("should return body if success", func() {
			ctrl := gomock.NewController(t)
			m, r, url, hash := initModule(ctrl)
			req, expiredAt, urlObject := genData()
			url.EXPECT().Create(m.db, req.Url, expiredAt).Return(urlObject, nil)
			hash.EXPECT().IDtoShortenID(urlObject.ID).Return("abc", nil)

			res := testUtils.FireRequest(r, http.MethodPost, path.ShortenApiPath, req)
			resp := &Response{}
			testUtils.VerifySuccessRes(res, http.StatusOK, resp)
			So(resp.Id, ShouldEqual, "abc")
			So(resp.ShortenUrl, ShouldEqual, "http://localhost:8080/abc")
		})
	})
}

func initModule(ctrl *gomock.Controller) (*Module, *gin.Engine, *MockIUrl, *MockIHash) {
	log := logrus.New()
	db, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{SkipDefaultTransaction: true})
	url := NewMockIUrl(ctrl)
	hash := NewMockIHash(ctrl)
	module := New(log, gormDB, url, hash)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST(path.ShortenApiPath, module.Controller)

	return module, r, url, hash
}

func genData() (*Request, time.Time, *database.Url) {
	req := &Request{
		Url:       "https://example.com",
		ExpiredAt: "2021-02-08T09:20:41Z",
	}

	expiredAt, _ := time.Parse("2006-01-02T15:04:05Z", req.ExpiredAt)

	url := &database.Url{
		ID:        10,
		Url:       req.Url,
		ExpiredAt: expiredAt,
	}
	return req, expiredAt, url
}
