package utils

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/nickhsine/test_backend/models"
	"gopkg.in/matryer/try.v1"

	log "github.com/Sirupsen/logrus"
)

// InitDB initiates the MySQL database connection
func InitDB(attempts, retryMaxDelay int) (*gorm.DB, error) {
	var db *gorm.DB
	err := try.Do(func(attempt int) (bool, error) {
		var err error

		// connect to MySQL database
		db, err = gorm.Open("mysql", Cfg.DBSettings.User+":"+Cfg.DBSettings.Password+"@tcp("+Cfg.DBSettings.Address+":"+Cfg.DBSettings.Port+")/"+Cfg.DBSettings.Name+"?parseTime=true")

		if err != nil {
			time.Sleep(time.Duration(retryMaxDelay) * time.Second)
		}

		return attempt < attempts, err
	})

	if err != nil {
		log.Error("Please check the MySQL database connection: ", err.Error())
		return nil, err
	}

	db.AutoMigrate(&models.Event{})
	//db.LogMode(true)

	return db, nil
}
