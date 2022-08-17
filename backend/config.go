package backend

import (
	"context"
	"my-playground/backend/model"
	"my-playground/backend/server"
	"my-playground/backend/tray"
	"my-playground/backend/utils"
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
	} else {
		config.loadOptions(options)
	}

	return config
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
		}
	}
}
