package main

import (
	"github.com/artemkaxboy/go-hocon"
	"github.com/sirupsen/logrus"
)

// Properties struct is used for loading and providing access to configuration file.
type Properties struct {
	Debug bool `hocon:"node=debug,default=false"`

	Server struct {
		Port int64 `hocon:"node=port,default=8080"`
	} `hocon:"node=server"`

	Filters struct {
		ExtremeValues struct {
			Enabled   bool    `hocon:"default=true"`
			Threshold float64 `hocon:"default=5"`
		}
	}

	Parser struct {
		URL string `hocon:"node=url"`
	} `hocon:"node=parser"`
}

var (
	props *Properties
)

// getProperties loads configuration from file to Properties struct if needed and gives pointer to it
func getProperties() *Properties {
	if props == nil {
		props = &Properties{}
		if err := hocon.LoadConfigFile("ticker-parser.conf", props); err != nil {
			logrus.WithError(err).Fatal("cannot load properties")
		}
	}
	return props
}
