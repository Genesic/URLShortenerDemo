package url

import "github.com/sirupsen/logrus"

type Module struct {
	log *logrus.Logger
}

func New(log *logrus.Logger) *Module {
	return &Module{
		log: log,
	}
}
