package model

import (
	"my-playground/backend/utils"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
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

	db.AutoMigrate(&MpConfig{})
}
