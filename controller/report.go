package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reportify-backend/entity"
)

type ReportController struct {
	ReportUseCase ReportUseCase
}

func NewReportController(u ReportUseCase) *ReportController {
	return &ReportController{u}
}

type GetReportsQuery struct {
	UserID string `form:"userId"`
}

// ReportResponse - 日報レスポンス
type ReportResponse struct {
	// 日報レスポンス
	Id string `json:"id"`
	// ユーザID
	UserId string `json:"userId"`
	// 本文
	Body string `json:"body"`
	// レビュー本文
	ReviewBody *string `json:"reviewBody"`
	// 実施したタスクリスト
	Tasks []entity.Task `json:"tasks"`
}

// ReportsResponse - 日報リストレスポンス
type ReportsResponse struct {
	// 日報リスト
	Reports []ReportResponse `json:"reports"`
}

// GetReports godoc
// @Summary 日報リスト取得API
// @Tags    Report
// @Accept  json
// @Produce json
// @Param   userId           query    string               true "ユーザID"
// @Param   organizationCode path     string               true "組織コード"
// @Success 200              {object} ReportsResponse      "OK"
// @Failure 400              {object} entity.ErrorResponse "BadRequest"
// @Failure 401              {object} entity.ErrorResponse "Unauthorized"
// @Failure 403              {object} entity.ErrorResponse "Forbidden"
// @Failure 404              {object} entity.ErrorResponse "Not Found"
// @Router  /organizations/{organizationCode}/reports [get]
func (c *ReportController) GetReports(ctx *gin.Context) (interface{}, error) {
	var query GetReportsQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	if query.UserID == "" {
		return nil, entity.NewError(http.StatusBadRequest, errors.New("userId is required"))
	}
	code := ctx.Params.ByName("organizationCode")
	fmt.Println(code)
	fmt.Println(query.UserID)
	panic("not implemented")
}

// GetReport godoc
// @Summary 日報取得API
// @Tags    Report
// @Accept  json
// @Produce json
// @Param   userId           query    string               true "ユーザID"
// @Param   organizationCode path     string               true "組織コード"
// @Param   organizationCode path     string               true "日報ID"
// @Success 200              {object} ReportResponse       "OK"
// @Failure 401              {object} entity.ErrorResponse "Unauthorized"
// @Failure 403              {object} entity.ErrorResponse "Forbidden"
// @Failure 404              {object} entity.ErrorResponse "Not Found"
// @Router  /organizations/{organizationCode}/reports/{reportId} [get]
func (c *ReportController) GetReport(ctx *gin.Context) (interface{}, error) {
	var req CreateReportRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	code := ctx.Params.ByName("organizationCode")
	fmt.Println(code)
	fmt.Println(req.Body)
	panic("not implemented")
}

// CreateReportRequest - 日報作成リクエスト
type CreateReportRequest struct {
	// 本文
	Body string `json:"body"`
	// 実施したタスクリスト
	Tasks []entity.Task `json:"tasks"`
}

// CreateReport godoc
// @Summary 日報作成API
// @Tags    Report
// @Accept  json
// @Produce json
// @Param   userId           query    string               true "ユーザID"
// @Param   organizationCode path     string               true "組織コード"
// @Success 201              {object} CreateReportRequest  "Created"
// @Failure 400              {object} entity.ErrorResponse "BadRequest"
// @Failure 401              {object} entity.ErrorResponse "Unauthorized"
// @Failure 403              {object} entity.ErrorResponse "Forbidden"
// @Failure 404              {object} entity.ErrorResponse "Not Found"
// @Router  /organizations/{organizationCode}/reports [post]
func (c *ReportController) CreateReport(ctx *gin.Context) (interface{}, error) {
	var req CreateReportRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	code := ctx.Params.ByName("organizationCode")
	fmt.Println(code)
	fmt.Println(req.Body)
	panic("not implemented")
}
