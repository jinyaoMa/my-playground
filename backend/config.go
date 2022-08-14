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
		Name:  "Server.HttpPort",
		Value: config.Server.HttpPort,
	})
	options = append(options, model.MpOption{
		Name:  "Server.HttpsPort",
		Value: config.Server.HttpsPort,
	})
	options = append(options, model.MpOption{
		Name:  "Server.CertsDirCache",
		Value: config.Server.CertsDirCache,
	})
	options = append(options, model.MpOption{
		Name:  "Tray.Language",
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
		case "Server.HttpPort":
			config.Server.HttpPort = option.Value
		case "Server.HttpsPort":
			config.Server.HttpsPort = option.Value
		case "Server.CertsDirCache":
			config.Server.CertsDirCache = option.Value
		case "Tray.Language":
			config.Tray.Language = option.Value
		}
	}
}
