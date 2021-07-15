package main

import (
	"URLShortenerDemo/service/database"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

var Schemas = []interface{}{
	&database.Url{},
}

func main() {
	log := logrus.New()
	db := database.New()
	if err := db.AutoMigrate(Schemas...); err != nil {
		log.WithField("err", err).Fatal("auto migrate schemas failed")
	}

	log.Info("auto migrate schemas success")
}
