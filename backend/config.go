package backend

import (
	"context"
	"log"
	"my-playground/backend/model"
	"my-playground/backend/server"
	"my-playground/backend/tray"
)

type Config struct {
	Server server.Config
	Tray   tray.Config
}

func LoadConfig(ctx context.Context) *Config {
	config := &Config{
		Server: server.Config{
			HttpPort:      ":10080",
			HttpsPort:     ":10433",
			CertsDirCache: "",
		},
		Tray: tray.Config{
			WailsCtx: ctx,
			Language: "zh",
		},
	}

	var options model.MpOptions
	result := options.Load()
	if result.Error != nil {
		log.Fatalln("fail to load options")
	}
	if result.RowsAffected == 0 {
		// options not yet generated and stored
		saveConfig2Options(config, options)
	} else {
		fillOptions2Config(options, config)
	}

	return config
}

func saveConfig2Options(config *Config, options model.MpOptions) {
	options = append(options, model.MpOption{
		Name:  server.CfgNameHttpPort,
		Value: config.Server.HttpPort,
	})
	options = append(options, model.MpOption{
		Name:  server.CfgNameHttpsPort,
		Value: config.Server.HttpsPort,
	})
	options = append(options, model.MpOption{
		Name:  server.CfgNameCertsDirCache,
		Value: config.Server.CertsDirCache,
	})
	options = append(options, model.MpOption{
		Name:  tray.CfgNameLanguage,
		Value: config.Tray.Language,
	})

	result := options.Save()
	if result.Error != nil {
		log.Fatalln("fail to save options")
	}
}

func fillOptions2Config(options model.MpOptions, config *Config) {
	for _, option := range options {
		switch option.Name {
		case server.CfgNameHttpPort:
			config.Server.HttpPort = option.Value
		case server.CfgNameHttpsPort:
			config.Server.HttpsPort = option.Value
		case server.CfgNameCertsDirCache:
			config.Server.CertsDirCache = option.Value
		case tray.CfgNameLanguage:
			config.Tray.Language = option.Value
		}
	}
}
