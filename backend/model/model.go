package model

import (
	"my-playground/backend/utils"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		&gorm.Config{
			FullSaveAssociations: false,
			Logger: logger.New(
				utils.Logger(PkgName),
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Error,
					IgnoreRecordNotFoundError: true,
					Colorful:                  utils.IsDev(),
				},
			),
			PrepareStmt: true,
		},
	)
	if err != nil {
		utils.Logger(PkgName).Fatalf("failed to connect database: %+v\n", err)
	}

	db.AutoMigrate(
		&MpOption{},
	)
}
