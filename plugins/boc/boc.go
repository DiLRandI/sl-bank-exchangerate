package main

import (
	"fmt"
	"os"

	"github.com/DiLRandI/sl-bank-exchangerate/app/plugins/boc"
	"github.com/DiLRandI/sl-bank-exchangerate/app/scraper"
	"github.com/DiLRandI/sl-bank-exchangerate/contract"
	"github.com/rs/zerolog"
)

const pluginName = "BOC"

var (
	bocLogger zerolog.Logger
	Version   = "devlopment"
)

func init() {
	bocLogger = zerolog.New(os.Stdout).With().
		Str("plugin", pluginName).
		Str("version", Version).
		Logger()
}

func Initialize(url string) contract.Convert {
	bocLogger.Info().Msgf("initializing with endpoint %q", url)

	scraper := scraper.New(&scraper.Config{
		URL: url,
	})
	lookup := boc.New(&boc.Config{
		Scraper: scraper,
		Logger:  bocLogger,
	})

	return func(value string) (int, error) {
		bocLogger.Info().Msgf("looking for value %s", value)
		val, err := lookup.Lookup(value)
		if err != nil {
			return 0, fmt.Errorf("error looking up value %s, %w", value, err)
		}

		return val, nil
	}
}
