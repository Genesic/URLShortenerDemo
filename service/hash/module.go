package hash

import (
	"URLShortenerDemo/pkg/env"
	"github.com/sirupsen/logrus"
	"github.com/speps/go-hashids"
)

type Module struct {
	log *logrus.Logger `inject:""`
	hd  *hashids.HashIDData
}

func New(log *logrus.Logger) *Module {
	hd := hashids.NewData()
	hd.Salt = env.GetDefault("hash_salt", "url_shortener_demo")

	return &Module{
		log: log,
		hd:  hd,
	}
}
