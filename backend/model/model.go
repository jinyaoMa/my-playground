package model

import (
	"log"
	"my-playground/backend/utils"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	path, err := utils.GetExecutablePath()
	if err != nil {
		log.Fatalln("failed to get executable path")
	}

	db, err = gorm.Open(
		sqlite.Open(filepath.Join(path, "my-playground.db")),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	db.AutoMigrate(&MpOption{})
}
