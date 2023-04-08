package boc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DiLRandI/sl-bank-exchangerate/app/scraper"
	"github.com/DiLRandI/sl-bank-exchangerate/common"
	"github.com/rs/zerolog"
)

type Lookup struct {
	scraper *scraper.Scraper
	logger  zerolog.Logger
}

type Config struct {
	Scraper *scraper.Scraper
	Logger  zerolog.Logger
}

func New(c *Config) *Lookup {
	return &Lookup{
		scraper: c.Scraper,
		logger:  c.Logger,
	}
}

func (l *Lookup) Lookup(code string) (int, error) {
	key, err := lookupMatcher(code)
	if err != nil {
		return 0, err
	}

	var convErr error

	value := 0
	keyFound := false

	if err := l.scraper.Visit("td", func(content string) {
		if keyFound {
			val, err := strconv.ParseFloat(strings.TrimSpace(content), 64)
			if err != nil {
				convErr = err
			}

			l.logger.Info().Msgf("found value for %s, %v", code, val)
			newVal := val * common.CentsFactor
			value = int(newVal)
			keyFound = !keyFound
		}

		if key == strings.TrimSpace(content) {
			l.logger.Info().Msgf("key %q is found", key)
			keyFound = !keyFound
		}
	}); err != nil {
		return 0, fmt.Errorf("error visiting, %w", err)
	}

	return value, convErr
}

func lookupMatcher(code string) (string, error) {
	for _, v := range common.LookupKeys {
		switch v {
		case "USD":
			return "US Dollar", nil
		case "EUR":
			return "Euro", nil
		}
	}

	return "", fmt.Errorf("key %q was not mapped, %w", code, ErrInvalidCode)
}
