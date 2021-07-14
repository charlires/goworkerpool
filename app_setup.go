package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func createLogger(logLevel string) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log_level": logLevel,
		}).Error("parsing log_level")

		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Out = os.Stdout
	// if cfg.Env != config.EnvDev {
	// 	logger.Formatter = &logrus.JSONFormatter{}
	// }
	return logger, nil
}
