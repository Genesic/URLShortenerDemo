package main

import (
	"URLShortenerDemo/service/database"
	"fmt"
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

	log.WithField("schemas", fmt.Sprintf("%+v", Schemas)).Info("auto migrate schemas success")
}
