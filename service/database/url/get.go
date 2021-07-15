package url

import (
	"URLShortenerDemo/pkg/errors"
	"URLShortenerDemo/service/database"
	oriError "errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (m *Module) Get(tx *gorm.DB, id uint) (*database.Url, *errors.ServiceError) {
	row := &database.Url{ID: id}
	if err := tx.First(row).Error; err != nil {
		if oriError.Is(err, gorm.ErrRecordNotFound) {
			m.log.WithField("id", id).Error("[url] not found error")
			return nil, errors.UrlNotFoundError
		}

		m.log.WithFields(logrus.Fields{
			"err": err,
			"id":  id,
		}).Error("[url] fetch url failed")
		return nil, errors.FetchDatabaseFailedError
	}

	return row, nil
}
