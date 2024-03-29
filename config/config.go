package config

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

var Config = getConfig()

type config struct {
	JWT_KEY string

	PASSWORD string
	USERNAME string
}

func getConfig() config {
	b, err := os.ReadFile("config.json")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to read config")
	}
	var c config
	err = json.Unmarshal(b, &c)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to unmarshal config")
	}
	return c
}
