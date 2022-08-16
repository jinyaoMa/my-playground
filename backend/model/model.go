package model

import (
	"my-playground/backend/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	PkgName = "model"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(
		sqlite.Open(utils.GetExecutablePath("my-playground.db")),
		&gorm.Config{},
	)
	if err != nil {
		utils.Logger(PkgName).Fatalf("failed to connect database: %+v\n", err)
	}

	db.AutoMigrate(
		&MpOption{},
	)
}
