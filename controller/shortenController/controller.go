package shortenController

import (
	"URLShortenerDemo/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (m *Module) Controller(c *gin.Context) {
	req, err := m.validateRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	expiredAt, err := m.parseExpiredAt(req.ExpiredAt)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	row, err := m.url.Create(m.db, req.Url, expiredAt)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	shortenId, err := m.hash.IDtoShortenID(row.ID)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	resp := &Response{
		Id:         shortenId,
		ShortenUrl: m.cfg.domain + shortenId,
	}
	c.JSON(http.StatusOK, resp)
}

func (m *Module) validateRequest(c *gin.Context) (*Request, *errors.ServiceError) {
	req := &Request{}
	if err := c.ShouldBindJSON(req); err != nil {
		m.log.WithField("err", err).Warn("[shortenController] validate request failed")
		return nil, errors.ValidateRequestFailedError
	}

	return req, nil
}

func (m *Module) parseExpiredAt(expiredAt string) (time.Time, *errors.ServiceError) {
	result, err := time.Parse("2006-01-02T15:04:05Z", expiredAt)
	if err != nil {
		m.log.WithFields(logrus.Fields{
			"err":       err,
			"expiredAt": expiredAt,
		}).Error("[shortenController] parse time failed")
		return time.Time{}, errors.ValidateRequestFailedError
	}

	return result, nil
}
