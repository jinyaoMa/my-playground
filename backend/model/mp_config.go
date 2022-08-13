package model

import "gorm.io/gorm"

type MpConfig struct {
	gorm.Model
	Name  string `gorm:"unique"` // Config setting name
	Value string ``              // Config setting value associated with name
}

type MpConfigs []MpConfig

func (mcs *MpConfigs) Load() (err error) {
	result := db.Find(mcs)
	return result.Error
}
