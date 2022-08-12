package model

type MpConfig struct {
}

type MpConfigs []MpConfig

func (mcs *MpConfigs) Load() (err error) {
	result := db.Find(mcs)
	return result.Error
}
