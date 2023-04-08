package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type ContentFn func(string)

type Config struct {
	URL string
}

type Scraper struct {
	url       string
	collector *colly.Collector
}

func New(c *Config) *Scraper {
	return &Scraper{
		url:       c.URL,
		collector: colly.NewCollector(colly.AllowURLRevisit(), colly.CacheDir("./temp")),
	}
}

func (s *Scraper) Register(tag string, fn ContentFn) {
	s.collector.OnHTML(tag, func(h *colly.HTMLElement) {
		fn(h.Text)
	})
}

func (s *Scraper) Visit() error {
	if err := s.collector.Visit(s.url); err != nil {
		return fmt.Errorf("unable to visit the url %s, %w", s.url, err)
	}

	return nil
}
