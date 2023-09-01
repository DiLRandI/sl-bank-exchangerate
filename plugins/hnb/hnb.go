package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/DiLRandI/sl-bank-exchangerate/app/plugins/hnb"
	"github.com/DiLRandI/sl-bank-exchangerate/app/scraper"
	"github.com/DiLRandI/sl-bank-exchangerate/contract"
)

const pluginName = "HNB"

var (
	hnbLogger *slog.Logger
	Version   = "devlopment"
)

func init() {
	hnbLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil)).
		With(pluginName, Version)
}

func Initialize(url string) contract.Convert {
	hnbLogger.Info("initializing with endpoint", url)

	scraper := scraper.New(&scraper.Config{
		URL: url,
	})
	lookup := hnb.New(&hnb.Config{
		Scraper: scraper,
		Logger:  hnbLogger,
	})

	return func(value string) (int, error) {
		hnbLogger.Info("looking for value", value)

		val, err := lookup.Lookup(value)
		if err != nil {
			return 0, fmt.Errorf("error looking up value %s, %w", value, err)
		}

		return val, nil
	}
}
