package config

import (
	"Moddormy_backend/utils/logger"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// C = new(config) -> C  *config
// C config -> C config -> no pointer

var C = &config{}

func init() {
	// Load configurations to struct
	yml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		logger.Log(logrus.Fatal, "UNABLE TO READ YAML CONFIGURATION FILE")
	}
	err = yaml.Unmarshal(yml, C)
	if err != nil {
		logger.Log(logrus.Fatal, "UNABLE TO PARSE YAML CONFIGURATION FILE")
	}

	// Apply configurations
	logrus.SetLevel(logrus.WarnLevel)
}
