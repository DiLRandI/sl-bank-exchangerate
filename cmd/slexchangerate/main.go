package main

import (
	"os"
	"plugin"

	"github.com/DiLRandI/sl-bank-exchangerate/config"
	"github.com/DiLRandI/sl-bank-exchangerate/contract"
	"github.com/rs/zerolog"
)

var Version = "development"

const appName = "SL Bank Exchange Rate Monitor"

func main() {
	logger := zerolog.New(os.Stdout).With().
		Str("app", appName).
		Str("version", Version).
		Logger()

	logger.Info().Msg("starting the app")

	config, err := config.ParseConfig("config.json")
	if err != nil {
		logger.Fatal().Err(err)
	}

	plugins := make(map[string]contract.PluginContract)

	for _, p := range config.Plugins {
		plugin, err := plugin.Open(p.File)
		if err != nil {
			logger.Err(err).Msgf("unable to open the plugin %s", p.Name)
		}

		fnLookup, err := plugin.Lookup("Convert")
		if err != nil {
			logger.Fatal().Err(err)
		}

		fn, ok := fnLookup.(contract.PluginContract)
		if !ok {
			logger.Fatal().Msg("function not satisfy the contract")
		}

		plugins[p.Name] = fn
	}
}
