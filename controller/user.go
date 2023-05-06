package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	Users []*entity.User `json:"users"`
}

// GetUsers godoc
// @Summary 組織のユーザリスト取得API
// @Tags    User
// @Accept  json
// @Produce json
// @Param   organizationCode path     string               true "組織コード"
// @Success 200              {object} UsersResponse        "OK"
// @Failure 401 {object} entity.ErrorResponse "Unauthorized"
// @Failure 403              {object} entity.ErrorResponse "Forbidden"
// @Failure 404 {object} entity.ErrorResponse "Not Found"
// @Router  /organizations/{organizationCode}/users [get]
func (c UserController) GetUsers(ctx *gin.Context) (interface{}, error) {
	code := ctx.Param("organizationCode")
	users, err := c.UserUseCase.GetUsers(ctx, code)
	return UsersResponse{users}, err
}

// InviteUserRequest - メンバー招待リクエスト
type InviteUserRequest struct {
	// メールアドレス
	Email string `json:"email"`
}

// InviteUser godoc
// @Summary メンバー招待API
// @Tags    User
// @Accept  json
// @Produce json
// @Param   organizationCode path     string               true "組織コード"
// @Param   request          body     InviteUserRequest    true "メンバー招待リクエスト"
// @Success 200              {object} entity.User          "OK"
// @Failure 400              {object} entity.ErrorResponse "BadRequest"
// @Failure 401              {object} entity.ErrorResponse "Unauthorized"
// @Failure 403              {object} entity.ErrorResponse "Forbidden"
// @Failure 409              {object} entity.ErrorResponse "Conflict"
// @Router  /organizations/{organizationCode}/users [post]
func (c UserController) InviteUser(ctx *gin.Context) (interface{}, error) {
	var req InviteUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	code := ctx.Param("organizationCode")
	userID, _ := ctx.Get(entity.ContextKeyUserID)

	return c.UserUseCase.InviteUser(ctx, req.Email, code, userID.(string))
}

// GetMe godoc
// @Summary ログインユーザー取得API
// @Tags    User
// @Accept  json
// @Produce json
// @Success 200 {object} entity.User          "OK"
// @Failure 401              {object} entity.ErrorResponse "Unauthorized"
// @Failure 404              {object} entity.ErrorResponse "Not Found"
// @Router  /users/me [get]
func (c UserController) GetMe(ctx *gin.Context) (interface{}, error) {
	userID, _ := ctx.Get(entity.ContextKeyUserID)
	return c.UserUseCase.GetUser(ctx, userID.(string))
}
