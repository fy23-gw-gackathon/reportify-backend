package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"reportify-backend/entity"
	"strings"
)

type UserRepo interface {
	GetUserIDFromToken(ctx context.Context, token string) (*string, error)
}

func Authentication(
	repo UserRepo,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerKey := c.Request.Header.Get("authorization")
		token := strings.Replace(bearerKey, "Bearer ", "", 1)
		userID, err := repo.GetUserIDFromToken(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entity.NewError(http.StatusUnauthorized, err))
			return
		}
		c.Set(entity.ContextKeyUserID, *userID)
		c.Set(entity.ContextKeyUserID, "test")
		c.Next()
	}
}
