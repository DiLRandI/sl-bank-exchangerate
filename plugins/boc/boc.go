package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

var bocLogger logrus.FieldLogger

func init() {
	ctx := context.Background()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"plugin":  "boc",
		"version": 0.1,
	})

	bocLogger = logger
}

func CanRun(msg string) {
	bocLogger.Info("Hi I got your message, ", msg)
}
