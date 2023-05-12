package middleware

import (
	"github.com/fy23-gw-gackathon/reportify-backend/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(cfg config.Config) gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowOrigins = cfg.AllowOrigins
	conf.AllowCredentials = true
	conf.AllowMethods = []string{
		"GET", "POST", "PUT", "DELETE",
	}
	conf.AllowHeaders = append(conf.AllowHeaders, "Authorization")
	return cors.New(conf)
}
