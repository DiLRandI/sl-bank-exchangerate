package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/DiLRandI/sl-bank-exchangerate/app/plugins/boc"
	"github.com/DiLRandI/sl-bank-exchangerate/app/scraper"
	"github.com/DiLRandI/sl-bank-exchangerate/contract"
)

const pluginName = "BOC"

var (
	bocLogger *slog.Logger
	Version   = "devlopment"
)

func init() {
	bocLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil)).
		With(pluginName, Version)
}

func Initialize(url string) contract.Convert {
	bocLogger.Info("initializing with endpoint", url)

	scraper := scraper.New(&scraper.Config{
		URL: url,
	})
	lookup := boc.New(&boc.Config{
		Scraper: scraper,
		Logger:  bocLogger,
	})

	return func(value string) (int, error) {
		bocLogger.Info("looking for value", value)
		val, err := lookup.Lookup(value)
		if err != nil {
			return 0, fmt.Errorf("error looking up value %s, %w", value, err)
		}

		return val, nil
	}
}
