package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"reportify-backend/config"
)

const TxKey = "transactionKey"
const ErrDuplicateEntryNumber = 1062

func NewDB(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.Database.Url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if cfg.App.Debug {
		db = db.Debug()
	}
	return db
}
