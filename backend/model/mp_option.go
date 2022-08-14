package model

import "gorm.io/gorm"

type MpOption struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"` // Option name
	Value string ``                   // Option value associated with name
}

func (mo *MpOption) Update(newValue string) *gorm.DB {
	return db.Model(mo).Where(mo).Updates(MpOption{
		Value: newValue,
	})
}

type MpOptions []MpOption

func (mos *MpOptions) Load() *gorm.DB {
	return db.Find(mos)
}

func (mos *MpOptions) Save() *gorm.DB {
	return db.Save(mos)
}
