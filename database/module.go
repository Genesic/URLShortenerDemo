package database

import (
	"URLShortenerDemo/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	dsn          string
	maxIdleConns int
	maxOpenConns int
}

func getConfig() *Config {
	return &Config{
		dsn:          env.Get("db_source"),
		maxIdleConns: env.GetDefaultInt("db_max_idle_conns", 5),
		maxOpenConns: env.GetDefaultInt("db_max_open_conns", 10),
	}
}

func New() *gorm.DB {
	cfg := getConfig()
	db, err := gorm.Open(mysql.Open(cfg.dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(cfg.maxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.maxOpenConns)
	return db
}
