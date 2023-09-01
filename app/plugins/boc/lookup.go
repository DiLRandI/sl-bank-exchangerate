package boc

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/DiLRandI/sl-bank-exchangerate/app/scraper"
	"github.com/DiLRandI/sl-bank-exchangerate/common"
)

type Lookup struct {
	scraper *scraper.Scraper
	logger  *slog.Logger
}

type Config struct {
	Scraper *scraper.Scraper
	Logger  *slog.Logger
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

	l.scraper.Register("td", func(content string) {
		if keyFound {
			val, err := strconv.ParseFloat(strings.TrimSpace(content), 64)
			if err != nil {
				convErr = err
			}

			l.logger.Info("found value for", code, val)
			newVal := val * common.CentsFactor
			value = int(newVal)
			keyFound = !keyFound
		}

		if key == strings.TrimSpace(content) {
			l.logger.Info("key= is found", key)
			keyFound = !keyFound
		}
	})

	if err := l.scraper.Visit(); err != nil {
		return 0, fmt.Errorf("error visiting, %w", err)
	}

	return value, convErr
}

func lookupMatcher(code string) (string, error) {
	switch code {
	case "USD":
		return "US Dollar", nil
	case "EUR":
		return "Euro", nil
	}

	return "", fmt.Errorf("key %q was not mapped, %w", code, ErrInvalidCode)
}
