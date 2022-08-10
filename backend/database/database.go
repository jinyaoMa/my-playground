package database

import (
	"my-playground/backend/utils"
	"path/filepath"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func Connect() *gorm.DB {
	once.Do(func() {
		path, err := utils.GetExecutablePath()
		if err != nil {
			panic("failed to get executable path")
		}

		db, err = gorm.Open(
			sqlite.Open(filepath.Join(path, "my-playground.db")),
			&gorm.Config{},
		)
		if err != nil {
			panic("failed to connect database")
		}
	})
	return db
}
