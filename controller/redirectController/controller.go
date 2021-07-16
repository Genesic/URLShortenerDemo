package redirectController

import (
	"URLShortenerDemo/pkg/errors"
	"URLShortenerDemo/service/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Request struct {
	ShortenId string `uri:"shorten_id" binding:"required"`
}

func (m *Module) Controller(c *gin.Context) {
	req, err := m.validateRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	id, err := m.hash.ShortenIDtoID(req.ShortenId)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	url, err := m.getUrl(id, req.ShortenId)
	if err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	if err := m.verifyExpired(url); err != nil {
		c.AbortWithStatusJSON(err.Status, err)
		return
	}

	c.Redirect(http.StatusFound, url.Url)
}

func (m *Module) validateRequest(c *gin.Context) (*Request, *errors.ServiceError) {
	req := &Request{}
	if err := c.ShouldBindUri(req); err != nil {
		m.log.WithField("err", err).Warn("[shortenController] validate request failed")
		return nil, errors.ValidateRequestFailedError
	}

	return req, nil
}

func (m *Module) getUrl(id uint, shortenId string) (*database.Url, *errors.ServiceError) {
	row, existed := m.cacheClient.Get(shortenId)
	if existed {
		if val, ok := row.(*database.Url); ok {
			return val, nil
		}
	}

	url, err := m.url.Get(m.db, id)
	if err != nil {
		return nil, err
	}

	m.cacheClient.Set(shortenId, url, m.cfg.expired)
	return url, nil
}

func (m *Module) verifyExpired(row *database.Url) *errors.ServiceError {
	if row.ExpiredAt.Before(time.Now()) {
		return errors.UrlNotFoundError
	}

	return nil
}
