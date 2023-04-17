package main

import (
	"fmt"
	"os"

	"github.com/DiLRandI/sl-bank-exchangerate/app/plugins/hnb"
	"github.com/DiLRandI/sl-bank-exchangerate/app/scraper"
	"github.com/DiLRandI/sl-bank-exchangerate/contract"
	"github.com/rs/zerolog"
)

const pluginName = "HNB"

var (
	hnbLogger zerolog.Logger
	Version   = "devlopment"
)

func init() {
	hnbLogger = zerolog.New(os.Stdout).With().
		Str("plugin", pluginName).
		Str("version", Version).
		Logger()
}

func Initialize(url string) contract.Convert {
	hnbLogger.Info().Msgf("initializing with endpoint %q", url)

	scraper := scraper.New(&scraper.Config{
		URL: url,
	})
	lookup := hnb.New(&hnb.Config{
		Scraper: scraper,
		Logger:  hnbLogger,
	})

	return func(value string) (int, error) {
		hnbLogger.Info().Msgf("looking for value %s", value)

		val, err := lookup.Lookup(value)
		if err != nil {
			return 0, fmt.Errorf("error looking up value %s, %w", value, err)
		}

		return val, nil
	}
}
