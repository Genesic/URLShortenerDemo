package url

import (
	"URLShortenerDemo/database"
	"URLShortenerDemo/errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func (m *Module) Create(tx *gorm.DB, url string, expiredAt time.Time) (*database.Url, *errors.ServiceError) {
	row := &database.Url{
		Url:       url,
		ExpiredAt: expiredAt,
	}
	if err := tx.Create(row).Error; err != nil {
		m.log.WithFields(logrus.Fields{
			"err": err,
			"row": fmt.Sprintf("%+v", row),
		}).Error("[url] create failed")
		return nil, errors.InsertDataBaseFailedError
	}

	return row, nil
}
