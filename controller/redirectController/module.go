package redirectController

import (
	"URLShortenerDemo/pkg/env"
	"URLShortenerDemo/pkg/errors"
	"URLShortenerDemo/service/database"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type Module struct {
	log         *logrus.Logger
	db          *gorm.DB
	cacheClient *cache.Cache
	hash        IHash
	url         IUrl
	cfg         *config
}

type config struct {
	expired time.Duration
}

func getConfig() *config {
	return &config{
		expired: time.Duration(env.GetDefaultInt("cache_expired", 3600)) * time.Second,
	}
}

type IHash interface {
	UrlIDtoID(shortenID string) (uint, *errors.ServiceError)
}

type IUrl interface {
	Get(tx *gorm.DB, id uint) (*database.Url, *errors.ServiceError)
}

func New(log *logrus.Logger, db *gorm.DB, cacheClient *cache.Cache, hash IHash, url IUrl) *Module {
	return &Module{
		log:         log,
		cfg:         getConfig(),
		db:          db,
		cacheClient: cacheClient,
		hash:        hash,
		url:         url,
	}
}
