package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reportify-backend/entity"
)

type OrganizationController struct {
	OrganizationUseCase
}

func NewOrganizationController(organizationUseCase OrganizationUseCase) *OrganizationController {
	return &OrganizationController{organizationUseCase}
}

type GetOrganizationsQuery struct {
	Limit  *int `form:"limit"`
	Offset *int `form:"offset"`
}

// UpdateOrganizationRequest - 組織更新リクエスト
type UpdateOrganizationRequest struct {
	// 組織名
	Name string `json:"name"`
	// 組織コード
	Code string `json:"code"`
	// ミッション
	Mission string `json:"mission"`
	// ビジョン
	Vision string `json:"vision"`
	// バリュー
	Value string `json:"value"`
}

// OrganizationsResponse - 組織リストレスポンス
type OrganizationsResponse struct {
	// 組織リスト
	Organizations []*entity.Organization `json:"organizations"`
}

// GetOrganizations godoc
// @Summary     組織リスト取得API
// @Description 自分が所属する組織のみ取得できる
// @Tags        Organization
// @Accept      json
// @Produce     json
// @Success     200 {object} OrganizationsResponse "OK"
// @Failure     400 {object} entity.ErrorResponse  "BadRequest"
// @Failure     401 {object} entity.ErrorResponse  "Unauthorized"
// @Router      /organizations [get]
// @Security    Bearer
func (c *OrganizationController) GetOrganizations(ctx *gin.Context) (interface{}, error) {
	userID, _ := ctx.Get(entity.ContextKeyUserID)
	id := userID.(string)
	o, err := c.OrganizationUseCase.GetOrganizations(ctx, id)
	return OrganizationsResponse{Organizations: o}, err
}

// GetOrganization godoc
// @Summary  組織取得API
// @Tags     Organization
// @Accept   json
// @Produce  json
// @Param    organizationCode path     string               true "組織コード"
// @Success  200              {object} entity.Organization  "OK"
// @Failure  400              {object} entity.ErrorResponse "BadRequest"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Router   /organizations/{organizationCode} [get]
// @Security Bearer
func (c *OrganizationController) GetOrganization(ctx *gin.Context) (interface{}, error) {
	userID, _ := ctx.Get(entity.ContextKeyUserID)
	id := userID.(string)
	code := ctx.Params.ByName("organizationCode")
	return c.OrganizationUseCase.GetOrganization(ctx, code, id)
}

// UpdateOrganization godoc
// @Summary  組織更新API
// @Tags     Organization
// @Accept   json
// @Produce  json
// @Param    organizationCode path     string                    true "組織コード"
// @Param    request          body     UpdateOrganizationRequest true "組織更新リクエスト"
// @Success  200              {object} entity.Organization       "OK"
// @Failure  400              {object} entity.ErrorResponse      "BadRequest"
// @Failure  401              {object} entity.ErrorResponse      "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse      "Forbidden"
// @Failure  404              {object} entity.ErrorResponse      "Not Found"
// @Failure  409              {object} entity.ErrorResponse      "Conflict"
// @Router   /organizations/{organizationCode} [put]
// @Security Bearer
func (c *OrganizationController) UpdateOrganization(ctx *gin.Context) (interface{}, error) {
	var req UpdateOrganizationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	userID, _ := ctx.Get(entity.ContextKeyUserID)
	id := userID.(string)
	code := ctx.Params.ByName("organizationCode")
	return c.OrganizationUseCase.UpdateOrganization(ctx, code, id, req.Name, req.Code, req.Mission, req.Vision, req.Value)
}
