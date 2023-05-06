package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"reportify-backend/config"
)

const TxKey = "transactionKey"
const ErrDuplicateEntryNumber = 1062

func NewDB(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.Database.Url), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if cfg.App.Env == "dev" {
		db = db.Debug()
	}
	return db
}
