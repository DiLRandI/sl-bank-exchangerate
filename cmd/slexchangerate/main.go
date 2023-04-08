package main

import (
	"context"
	"plugin"

	"github.com/DiLRandI/sl-bank-exchange-rate-console.git/config"
	"github.com/sirupsen/logrus"
)

const Version = "development"

const appName = "SL Bank Exchange Rate Monitor"

func main() {
	ctx := context.Background()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"app":     appName,
		"version": Version,
	})

	logger.Info("starting the app")

	config, err := config.ParseConfig("config.json")
	if err != nil {
		logger.Fatal(err)
	}

	for _, p := range config.Plugins {
		plugin, err := plugin.Open(p.File)
		if err != nil {
			logger.Errorf("unable to open the plugin %s, %v", p.Name, err)
		}

		fn, err := plugin.Lookup("CanRun")
		if err != nil {
			logrus.Fatal(err)
		}

		hello, ok := fn.(func(msg string))
		if !ok {
			logrus.Fatal("failed")
		}

		hello("Hello from main")
	}
}
