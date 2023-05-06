package middleware

import (
	"errors"
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
		if err != nil || userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, entity.NewError(http.StatusUnauthorized, errors.New("invalid token")))
			return
		}
		c.Set(entity.ContextKeyUserID, *userID)
		c.Next()
	}
}
