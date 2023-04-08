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
		Str("versoin", Version).
		Logger()
}

func CanRun(msg string) {
	bocLogger.Info().Msgf("Hi I got your message, ", msg)
}
