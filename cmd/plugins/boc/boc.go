package main

import (
	"os"

	"github.com/rs/zerolog"
)

const pluginName = "BOC"

var bocLogger zerolog.Logger
var Version = "devlopment"

func init() {
	bocLogger = zerolog.New(os.Stdout).With().
		Str("plugin", pluginName).
		Str("version", Version).
		Logger()
}

func Convert(string, string) (error, int) {
	return nil, 0
}
