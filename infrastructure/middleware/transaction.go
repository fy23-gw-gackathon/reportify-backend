package middleware

import (
	"github.com/fy23-gw-gackathon/reportify-backend/infrastructure/driver"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "gorm.io/gorm"

func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := db.Begin()
		defer func() {
			if http.StatusBadRequest <= c.Writer.Status() {
				tx.Rollback()
				return
			}
			tx.Commit()
		}()
		c.Set(driver.TxKey, tx)
		c.Next()
	}
}
