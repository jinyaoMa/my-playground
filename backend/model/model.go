package model

import (
	"log"
	"my-playground/backend/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(
		sqlite.Open(utils.GetExecutablePath("my-playground.db")),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	db.AutoMigrate(&MpOption{})
}
