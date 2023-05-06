package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	Organizations []entity.Organization `json:"organizations"`
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
func (c *OrganizationController) GetOrganizations(ctx *gin.Context) (interface{}, error) {
	panic("not implemented")
}

// GetOrganization godoc
// @Summary 組織取得API
// @Tags    Organization
// @Accept  json
// @Produce json
// @Param   organizationCode path     string               true "組織コード"
// @Success 200              {object} entity.Organization  "OK"
// @Failure 400              {object} entity.ErrorResponse      "BadRequest"
// @Failure 401              {object} entity.ErrorResponse      "Unauthorized"
// @Failure 403              {object} entity.ErrorResponse      "Forbidden"
// @Failure 404              {object} entity.ErrorResponse      "Not Found"
// @Router  /organizations/{organizationCode} [get]
func (c *OrganizationController) GetOrganization(ctx *gin.Context) (interface{}, error) {
	code := ctx.Params.ByName("organizationCode")
	fmt.Println(code)
	panic("not implemented")
}

// UpdateOrganization godoc
// @Summary 組織更新API
// @Tags    Organization
// @Accept  json
// @Produce json
// @Param   organizationCode path     string                    true "組織コード"
// @Param   request          body     UpdateOrganizationRequest true "組織更新リクエスト"
// @Success 200              {object} entity.Organization       "OK"
// @Failure 400              {object} entity.ErrorResponse "BadRequest"
// @Failure 401              {object} entity.ErrorResponse "Unauthorized"
// @Failure 403              {object} entity.ErrorResponse "Forbidden"
// @Failure 404              {object} entity.ErrorResponse "Not Found"
// @Failure 409              {object} entity.ErrorResponse      "Conflict"
// @Router  /organizations/{organizationCode} [put]
func (c *OrganizationController) UpdateOrganization(ctx *gin.Context) (interface{}, error) {
	code := ctx.Params.ByName("organizationCode")
	fmt.Println(code)
	panic("not implemented")
}
