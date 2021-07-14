package url

import (
	"URLShortenerDemo/database"
	"URLShortenerDemo/errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestModule_Create(t *testing.T) {
	Convey("test create method", t, func() {
		Convey("should return InsertDataBaseFailedError if insert failed", func() {
			m, tx, mock := initModule()
			url := genData()
			mock.ExpectExec("INSERT INTO `urls` (.*) VALUES (.*)").
				WithArgs(url.Url, url.ExpiredAt).
				WillReturnError(gorm.ErrInvalidData)

			result, err := m.Create(tx, url.Url, url.ExpiredAt)
			So(result, ShouldEqual, nil)
			So(err, ShouldEqual, errors.InsertDataBaseFailedError)
		})

		Convey("should return url object if success", func() {
			m, tx, mock := initModule()
			url := genData()
			mock.ExpectExec("INSERT INTO `urls` (.*) VALUES (.*)").
				WithArgs(url.Url, url.ExpiredAt).
				WillReturnResult(sqlmock.NewResult(12, 1))

			result, err := m.Create(tx, url.Url, url.ExpiredAt)
			So(result.ID, ShouldEqual, 12)
			So(result.Url, ShouldEqual, url.Url)
			So(result.ExpiredAt, ShouldEqual, url.ExpiredAt)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestModule_Get(t *testing.T) {
	Convey("test get method", t, func() {
		Convey("should return UrlNotFoundError if url not found or expired", func() {
			m, tx, mock := initModule()
			url := genData()
			mock.ExpectQuery("SELECT (.*) FROM `urls` WHERE expired_at > (.*) AND `urls`.`id` = (.*) ORDER BY `urls`.`id` LIMIT 1").
				WithArgs(sqlmock.AnyArg(), url.ID).WillReturnError(gorm.ErrRecordNotFound)

			result, err := m.Get(tx, url.ID)
			So(err, ShouldEqual, errors.UrlNotFoundError)
			So(result, ShouldEqual, nil)
		})

		Convey("should return FetchDatabaseFailedError if fetch url failed", func() {
			m, tx, mock := initModule()
			url := genData()
			mock.ExpectQuery("SELECT (.*) FROM `urls` WHERE expired_at > (.*) AND `urls`.`id` = (.*) ORDER BY `urls`.`id` LIMIT 1").
				WithArgs(sqlmock.AnyArg(), url.ID).WillReturnError(gorm.ErrInvalidData)

			result, err := m.Get(tx, url.ID)
			So(err, ShouldEqual, errors.FetchDatabaseFailedError)
			So(result, ShouldEqual, nil)
		})

		Convey("should return url object if success", func() {
			m, tx, mock := initModule()
			url := genData()
			row := mockUrl(url)
			mock.ExpectQuery("SELECT (.*) FROM `urls` WHERE expired_at > (.*) AND `urls`.`id` = (.*) ORDER BY `urls`.`id` LIMIT 1").
				WithArgs(sqlmock.AnyArg(), url.ID).WillReturnRows(row)

			result, err := m.Get(tx, url.ID)
			So(err, ShouldEqual, nil)
			So(result.ID, ShouldEqual, url.ID)
			So(result.Url, ShouldEqual, url.Url)
			So(result.ExpiredAt, ShouldEqual, url.ExpiredAt)
		})
	})
}

func initModule() (*Module, *gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{SkipDefaultTransaction: true})
	log := logrus.New()
	module := New(log)
	return module, gormDB, mock
}

func genData() *database.Url {
	row := &database.Url{
		ID:        123,
		Url:       "https://example.com",
		ExpiredAt: time.Now().AddDate(0, 0, 1),
	}

	return row
}

func mockUrl(url *database.Url) *sqlmock.Rows {
	row := sqlmock.NewRows(database.UrlColumns)
	row.AddRow(url.ID, url.Url, url.ExpiredAt)
	return row
}
