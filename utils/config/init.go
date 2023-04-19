package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var C = new(config)

// C = new(config) -> C  *config
// C config -> C config -> no pointer

func init() {
	// Load yaml configuration file to struct
	yml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		logrus.Fatal("UNABLE TO READ YAML CONFIGURATION FILE")
	}
	err = yaml.Unmarshal(yml, C)
	// C -> must be pointer
	if err != nil {
		logrus.Fatal("UNABLE TO PARSE YAML CONFIGURATION FILE")
	}

	// Apply log level configuration
	logrus.SetLevel(logrus.Level(C.LogLevel))
}
