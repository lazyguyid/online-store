package config

import (
	"errors"
	"online-store/core"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	config   core.Config
	postgres *gorm.DB
}

func NewStorage(app core.App) core.Storage {
	s := new(storage)
	s.config = app.Config()

	return s
}

func (s *storage) Postgres() (db *gorm.DB) {
	if s.postgres == nil {
		dsn := s.config.Get("DATABASE_URL").(string)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		s.postgres = db
	}

	return s.postgres
}

func (s *storage) Begin(engine string) (db *gorm.DB) {
	switch engine {
	case core.StorageEngines.Postgres:
		return s.Postgres().Begin()
	}

	panic(errors.New("invalid storage engine"))
}
