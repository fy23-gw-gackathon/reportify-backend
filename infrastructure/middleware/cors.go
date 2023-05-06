package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"reportify-backend/config"
)

func Cors(cfg config.Config) gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowOrigins = cfg.AllowOrigins
	conf.AllowCredentials = true
	return cors.New(conf)
}
