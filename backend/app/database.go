package app

import (
	"my-playground/backend/utils"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDB() (db *gorm.DB) {
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
	return
}
