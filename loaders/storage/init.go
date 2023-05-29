package storage

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"Moddormy_backend/utils/logger"
)

var RootDir string

func Init() {
	// Convert directory to absolute path
	if dir, err := filepath.Abs("./images"); err != nil {
		logger.Log(logrus.Fatal, "UNKNOWN STORAGE PATH")
	} else {
		RootDir = dir
	}

	// Confirm directory is existed
	if _, err := os.Stat(RootDir); os.IsNotExist(err) {
		logger.Log(logrus.Fatal, "NONEXISTENT STORAGE PATH")
	}
}
