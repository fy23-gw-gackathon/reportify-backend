package controller

import (
	"errors"
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	Users []*entity.User `json:"users"`
}

// GetUsers godoc
// @Summary  組織のユーザリスト取得API
// @Tags     User
// @Accept   json
// @Produce  json
// @Param    organizationCode path string                true "組織コード"
// @Success  200              {object} UsersResponse        "OK"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Router   /organizations/{organizationCode}/users [get]
// @Security Bearer
func (c UserController) GetUsers(ctx *gin.Context) (interface{}, error) {
	user, _ := ctx.Get(entity.ContextKeyUser)
	oUser := user.(*entity.OrganizationUser)
	users, err := c.UserUseCase.GetUsers(ctx, oUser.OrganizationID)
	return UsersResponse{users}, err
}

// InviteUserRequest - メンバー招待リクエスト
type InviteUserRequest struct {
	// メールアドレス
	Email string `json:"email"`
}

// InviteUser godoc
// @Summary  メンバー招待API
// @Tags     User
// @Accept   json
// @Produce  json
// @Param    organizationCode path string true "組織コード"
// @Param    request          body     InviteUserRequest    true "メンバー招待リクエスト"
// @Success  200              {object} entity.User          "OK"
// @Failure  400              {object} entity.ErrorResponse "BadRequest"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  409              {object} entity.ErrorResponse "Conflict"
// @Router   /organizations/{organizationCode}/users [post]
// @Security Bearer
func (c UserController) InviteUser(ctx *gin.Context) (interface{}, error) {
	var req InviteUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	user, _ := ctx.Get(entity.ContextKeyUser)
	oUser := user.(*entity.OrganizationUser)

	if !oUser.IsAdmin {
		return nil, entity.NewError(http.StatusForbidden, errors.New("you are not admin"))
	}

	return c.UserUseCase.InviteUser(ctx, req.Email, oUser.OrganizationID)
}

// GetMe godoc
// @Summary  ログインユーザー取得API
// @Tags     User
// @Accept   json
// @Produce  json
// @Success  200 {object} entity.User          "OK"
// @Failure  401 {object} entity.ErrorResponse "Unauthorized"
// @Failure  404 {object} entity.ErrorResponse "Not Found"
// @Router   /users/me [get]
// @Security Bearer
func (c UserController) GetMe(ctx *gin.Context) (interface{}, error) {
	bearerKey := ctx.Request.Header.Get("authorization")
	token := strings.Replace(bearerKey, "Bearer ", "", 1)
	return c.UserUseCase.GetUserFromToken(ctx, token)
}

// UpdateUserRoleRequest - ユーザーロール更新リクエスト
type UpdateUserRoleRequest struct {
	// ロール
	Role bool `json:"role"`
}

// UpdateUserRole godoc
// @Summary  ユーザーロール更新API
// @Tags     User
// @Accept   json
// @Produce  json
// @Param    organizationCode path     string               true "組織コード"
// @Param    userId           path string                true "ユーザーID"
// @Param    request          body UpdateUserRoleRequest true "ユーザーロール更新リクエスト"
// @Success  200              "OK"
// @Failure  400              {object} entity.ErrorResponse "BadRequest"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Failure  409              {object} entity.ErrorResponse "Conflict"
// @Router   /organizations/{organizationCode}/users/{userId} [put]
// @Security Bearer
func (c UserController) UpdateUserRole(ctx *gin.Context) (interface{}, error) {
	var req UpdateUserRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	user, _ := ctx.Get(entity.ContextKeyUser)
	oUser := user.(*entity.OrganizationUser)

	if !oUser.IsAdmin {
		return nil, entity.NewError(http.StatusForbidden, errors.New("you are not admin"))
	}

	userID := ctx.Param("userId")
	return nil, c.UserUseCase.UpdateUserRole(ctx, oUser.OrganizationID, userID, req.Role)
}

// DeleteUser godoc
// @Summary  ユーザー削除API
// @Tags     User
// @Accept   json
// @Produce  json
// @Param    organizationCode path     string               true "組織コード"
// @Param    userId           path string true "ユーザーID"
// @Success  200              "OK"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Router   /organizations/{organizationCode}/users/{userId} [delete]
// @Security Bearer
func (c UserController) DeleteUser(ctx *gin.Context) (interface{}, error) {
	user, _ := ctx.Get(entity.ContextKeyUser)
	oUser := user.(*entity.OrganizationUser)

	if !oUser.IsAdmin {
		return nil, entity.NewError(http.StatusForbidden, errors.New("you are not admin"))
	}

	userID := ctx.Param("userId")
	return nil, c.UserUseCase.DeleteUser(ctx, oUser.OrganizationID, userID)
}
