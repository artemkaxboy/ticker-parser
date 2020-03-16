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
		Catalog struct {
			BaseUrl  string `hocon:"node=baseUrl,default=www"`
			PageSize int16  `hocon:"node=pageSize,default=25"`
		} `hocon:"node=catalog"`

		URL string `hocon:"node=url,default=www.site.com"`
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
			logrus.Warn("cannot load properties file ticker-parser.conf")
			if err := hocon.LoadConfigText("", props); err != nil {
				logrus.WithError(err).Fatal("cannot load properties")
			}
		}
	}
	return props
}
