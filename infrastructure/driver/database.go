package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const TxKey = "transactionKey"

var (
	dsn    = os.Getenv("DATABASE_URL")
	appEnv = os.Getenv("APP_ENV")
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if appEnv == "dev" {
		db = db.Debug()
	}
	return db
}
