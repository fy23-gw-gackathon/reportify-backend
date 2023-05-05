package controller

import (
	"github.com/gin-gonic/gin"
	"reportify-backend/entity"
)

type UserController struct {
	UserUseCase
}

func NewUserController(u UserUseCase) *UserController {
	return &UserController{u}
}

// UsersResponse - ユーザリストレスポンス
type UsersResponse struct {
	// ユーザリスト
	Users []entity.User `json:"users"`
}

// GetUsers godoc
// @Summary 組織のユーザリスト取得API
// @Tags    User
// @Accept  json
// @Produce json
// @Param   organizationCode path     string               true "組織コード"
// @Success 200              {object} UsersResponse        "OK"
// @Failure 401              {object} entity.ErrorResponse "Unauthorized"
// @Failure 403              {object} entity.ErrorResponse "Forbidden"
// @Failure 404              {object} entity.ErrorResponse "Not Found"
// @Router  /organizations/{organizationCode}/users [get]
func (c UserController) GetUsers(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}
