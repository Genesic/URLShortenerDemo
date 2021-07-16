package main

import (
	"URLShortenerDemo/controller/redirectController"
	"URLShortenerDemo/controller/shortenController"
	"URLShortenerDemo/pkg/env"
	"URLShortenerDemo/pkg/path"
	"URLShortenerDemo/service/database"
	"URLShortenerDemo/service/database/url"
	"URLShortenerDemo/service/hash"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"time"
)

func main() {
	s := NewService()
	gin.SetMode(s.mode)
	r := gin.New()
	r.Use(ginlogrus.Logger(s.log), gin.Recovery())
	r.GET(path.RedirectApiPath, s.redirect.Controller)
	r.POST(path.ShortenApiPath, s.shorten.Controller)

	s.log.WithField("port", s.port).Info("server port")
	if err := r.Run(":" + s.port); err != nil {
		panic(err)
	}
}

type Service struct {
	port     string
	mode     string
	log      *logrus.Logger
	shorten  *shortenController.Module
	redirect *redirectController.Module
}

func NewService() *Service {
	log := logrus.New()
	db := database.New()

	cacheClient := cache.New(time.Hour, time.Minute)
	hashModule := hash.New(log)
	urlModule := url.New(log)

	shorten := shortenController.New(log, db, urlModule, hashModule)
	redirect := redirectController.New(log, db, cacheClient, hashModule, urlModule)

	return &Service{
		port:     env.GetDefault("port", "8080"),
		mode:     env.GetDefault("mode", "debug"),
		log:      log,
		shorten:  shorten,
		redirect: redirect,
	}
}