package hash

import (
	"URLShortenerDemo/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/speps/go-hashids"
)

func (m *Module) IDtoUrlID(id uint) (string, *errors.ServiceError) {
	hashID, _ := hashids.NewWithData(m.hd)
	value, err := hashID.EncodeInt64([]int64{int64(id)})
	if err != nil {
		m.log.WithFields(logrus.Fields{
			"id":  id,
			"err": err,
		}).Error("[hash] encode failed")
		return "", errors.EncodeHashFailedError
	}
	return value, nil
}

func (m *Module) UrlIDtoID(urlId string) (uint, *errors.ServiceError) {
	hashID, _ := hashids.NewWithData(m.hd)
	numbers, err := hashID.DecodeInt64WithError(urlId)
	if err != nil {
		m.log.WithFields(logrus.Fields{
			"urlId": urlId,
			"err":   err,
		}).Error("[hash] decode failed")
		return 0, errors.UrlNotFoundError
	} else if len(numbers) <= 0 {
		m.log.WithFields(logrus.Fields{
			"urlId": urlId,
			"err":   err,
		}).Error("[hash] decode empty result")
		return 0, errors.UrlNotFoundError
	}
	return uint(numbers[0]), nil
}
