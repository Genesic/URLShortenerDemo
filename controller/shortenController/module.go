package shortenController

import (
	"URLShortenerDemo/pkg/env"
	"URLShortenerDemo/pkg/errors"
	"URLShortenerDemo/service/database"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type Module struct {
	log  *logrus.Logger
	cfg  *config
	db   *gorm.DB
	url  IUrl
	hash IHash
}

type config struct {
	domain string
}

func getConfig() *config {
	return &config{
		domain: env.GetDefault("domain", "http://localhost:8080/"),
	}
}

type IUrl interface {
	Create(tx *gorm.DB, url string, expiredAt time.Time) (*database.Url, *errors.ServiceError)
}

type IHash interface {
	IDtoUrlID(id uint) (string, *errors.ServiceError)
}

func New(log *logrus.Logger, db *gorm.DB, url IUrl, hash IHash) *Module {
	return &Module{
		log:  log,
		cfg:  getConfig(),
		db:   db,
		url:  url,
		hash: hash,
	}
}
