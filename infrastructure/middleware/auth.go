package middleware

import (
	"github.com/fy23-gw-gackathon/reportify-backend/config"
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

type UserRepo interface {
	GetUserIDFromToken(ctx context.Context, token string) (*string, error)
	GetOrganizationUser(ctx context.Context, organizationCode string, userID string) (*entity.OrganizationUser, error)
}

func Authentication(
	repo UserRepo,
	cfg config.Config,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerKey := c.Request.Header.Get("authorization")
		token := strings.Replace(bearerKey, "Bearer ", "", 1)
		code := c.Params.ByName("organizationCode")
		var userID string
		// Debugモードの場合はユーザーIDを固定
		if cfg.App.Debug {
			userID = "01GZT0HJAX3CM9P1V66T9GYM95"
		} else {
			uid, err := repo.GetUserIDFromToken(context.Background(), token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, entity.ErrorResponse{Message: err.Error()})
				return
			}
			userID = *uid
		}

		var user *entity.OrganizationUser
		if code != "" {
			var err error
			user, err = repo.GetOrganizationUser(c, code, userID)
			if err != nil {
				c.JSON(http.StatusUnauthorized, entity.ErrorResponse{Message: err.Error()})
				return
			}
		} else {
			user = &entity.OrganizationUser{
				UserID: userID,
			}
		}
		c.Set(entity.ContextKeyUser, user)
		c.Next()
	}
}
