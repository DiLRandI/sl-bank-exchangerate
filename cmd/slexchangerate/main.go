package main

import (
	"os"
	"plugin"

	"github.com/DiLRandI/sl-bank-exchange-rate-console.git/config"
	"github.com/rs/zerolog"
)

var Version = "development"

const appName = "SL Bank Exchange Rate Monitor"

func main() {
	logger := zerolog.New(os.Stdout).With().
		Str("app", appName).
		Str("versoin", Version).
		Logger()

	logger.Info().Msg("starting the app")

	config, err := config.ParseConfig("config.json")
	if err != nil {
		logger.Fatal().Err(err)
	}

	for _, p := range config.Plugins {
		plugin, err := plugin.Open(p.File)
		if err != nil {
			logger.Err(err).Msgf("unable to open the plugin %s", p.Name)
		}

		fn, err := plugin.Lookup("CanRun")
		if err != nil {
			logger.Fatal().Err(err)
		}

		hello, ok := fn.(func(msg string))
		if !ok {
			logger.Fatal().Msg("unable to load the function")
		}

		hello("Hello from main")
	}
}
