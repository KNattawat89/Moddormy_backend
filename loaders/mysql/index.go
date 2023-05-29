package mysql

import (
	"Moddormy_backend/loaders/mysql/model"
	"Moddormy_backend/utils/config"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Gorm *gorm.DB

func Init() {
	// due to GORM log level have only 4 level but logrus have 7 level,
	// we need to pass from 4 level to 7 level
	gormLogLevel := []logger.LogLevel{
		logger.Silent,
		logger.Error,
		logger.Error,
		logger.Warn,
		logger.Info,
		logger.Info,
		logger.Info,
	}

	gormLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  gormLogLevel[config.C.LogLevel],
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		})

	// open SQL connection
	dialector := mysql.New(
		mysql.Config{
			DSN: config.C.MySqlDsn,
		},
	)

	if db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
		//DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		logrus.Fatal("UNABLE TO LOAD GORM MYSQL DATABASE")
	} else {
		Gorm = db
	}

	// Initialize model migrations
	if config.C.MySqlMigrate {
		if err := migrate(); err != nil {
			logrus.Fatal("UNABLE TO MIGRATE GORM MODEL")
		}
	}
	assignModel()
	logrus.Debugln("INITIALIZE MYSQL CONNECTION")
}
func migrate() error {
	if err := Gorm.AutoMigrate(
		new(model.DormImage),
		new(model.Dorm),
		new(model.Favorite),
		new(model.Review),
		new(model.RoomImage),
		new(model.Room),
		new(model.User),
	); err != nil {
		return err
	}
	return nil
}
