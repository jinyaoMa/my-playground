package backend

import (
	"context"
	"my-playground/backend/model"
	"my-playground/backend/server"
	"my-playground/backend/tray"
	"my-playground/backend/utils"
)

const (
	CfgTotalNumberOfOptions = 5
)

type Config struct {
	Server *server.Config
	Tray   *tray.Config
}

func LoadConfig(ctx context.Context) *Config {
	config := &Config{
		Server: &server.Config{
			HttpPort:      ":10080",
			HttpsPort:     ":10433",
			CertsDirCache: "",
		},
		Tray: &tray.Config{
			Context:  ctx,
			Language: "zh",
			Theme:    "light",
		},
	}

	var options model.MpOptions
	result := options.Load()
	if result.Error != nil {
		utils.Logger(PkgName).Fatalf("fail to load options: %+v\n", result.Error)
	}
	if result.RowsAffected == 0 {
		// options not yet generated and stored
		config.saveOptions(options)
	} else if result.RowsAffected < CfgTotalNumberOfOptions {
		config.loadOptions(options)
		config.updateOptions(options)
	} else {
		config.loadOptions(options)
	}

	return config
}

func (c *Config) updateOptions(options model.MpOptions) {
	optionPairs := [][]string{
		{server.CfgNameHttpPort, c.Server.HttpPort},
		{server.CfgNameHttpsPort, c.Server.HttpsPort},
		{server.CfgNameCertsDirCache, c.Server.CertsDirCache},
		{tray.CfgNameLanguage, c.Tray.Language},
		{tray.CfgNameTheme, c.Tray.Theme},
	}

	var newOptions model.MpOptions
	for _, pair := range optionPairs {
		optionNotExist := true
		for _, option := range options {
			// update exist options
			if option.Name == pair[0] {
				option.Value = pair[1]
				optionNotExist = false
				break
			}
		}
		if optionNotExist {
			// initialize new options
			newOptions = append(newOptions, model.MpOption{
				Name:  pair[0],
				Value: pair[1],
			})
		}
	}

	options = append(options, newOptions...)

	result := options.Save()
	if result.Error != nil {
		utils.Logger(PkgName).Fatalf("fail to update options: %+v\n", result.Error)
	}
}

func (c *Config) saveOptions(options model.MpOptions) {
	options = append(options, model.MpOption{
		Name:  server.CfgNameHttpPort,
		Value: c.Server.HttpPort,
	})
	options = append(options, model.MpOption{
		Name:  server.CfgNameHttpsPort,
		Value: c.Server.HttpsPort,
	})
	options = append(options, model.MpOption{
		Name:  server.CfgNameCertsDirCache,
		Value: c.Server.CertsDirCache,
	})
	options = append(options, model.MpOption{
		Name:  tray.CfgNameLanguage,
		Value: c.Tray.Language,
	})
	options = append(options, model.MpOption{
		Name:  tray.CfgNameTheme,
		Value: c.Tray.Theme,
	})

	result := options.Save()
	if result.Error != nil {
		utils.Logger(PkgName).Fatalf("fail to save options: %+v\n", result.Error)
	}
}

func (c *Config) loadOptions(options model.MpOptions) {
	for _, option := range options {
		switch option.Name {
		case server.CfgNameHttpPort:
			c.Server.HttpPort = option.Value
		case server.CfgNameHttpsPort:
			c.Server.HttpsPort = option.Value
		case server.CfgNameCertsDirCache:
			c.Server.CertsDirCache = option.Value
		case tray.CfgNameLanguage:
			c.Tray.Language = option.Value
		case tray.CfgNameTheme:
			c.Tray.Theme = option.Value
		}
	}
}
