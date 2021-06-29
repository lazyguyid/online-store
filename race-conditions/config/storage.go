package config

import (
	"errors"

	"github.com/lazyguyid/gacor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	config   gacor.Config
	postgres *gorm.DB
}

func NewStorage(app gacor.App) gacor.Storage {
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
	case gacor.StorageEngines.Postgres:
		return s.Postgres().Begin()
	}

	panic(errors.New("invalid storage engine"))
}
