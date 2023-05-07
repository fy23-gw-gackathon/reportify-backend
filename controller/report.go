package controller

import (
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

// ReportsResponse - 日報リストレスポンス
type ReportsResponse struct {
	// 日報リスト
	Reports []*entity.Report `json:"reports"`
}

// GetReports godoc
// @Summary  日報リスト取得API
// @Tags     Report
// @Accept   json
// @Produce  json
// @Param    organizationCode path     string               true "組織コード"
// @Success  200              {object} ReportsResponse      "OK"
// @Failure  400              {object} entity.ErrorResponse "BadRequest"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Router   /organizations/{organizationCode}/reports [get]
// @Security Bearer
func (c *ReportController) GetReports(ctx *gin.Context) (interface{}, error) {
	userID, _ := ctx.Get(entity.ContextKeyUserID)
	id := userID.(string)
	code := ctx.Params.ByName("organizationCode")
	fmt.Println(code)
	fmt.Println(id)
	reports, err := c.ReportUseCase.GetReports(ctx, code, id)
	return ReportsResponse{reports}, err
}

// GetReport godoc
// @Summary  日報取得API
// @Tags     Report
// @Accept   json
// @Produce  json
// @Param    userId           query    string               true "ユーザID"
// @Param    organizationCode path     string               true "組織コード"
// @Param    reportId         path     string               true "日報ID"
// @Success  200              {object} entity.Report        "OK"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Router   /organizations/{organizationCode}/reports/{reportId} [get]
// @Security Bearer
func (c *ReportController) GetReport(ctx *gin.Context) (interface{}, error) {
	code := ctx.Params.ByName("organizationCode")
	reportId := ctx.Params.ByName("reportId")
	userID, _ := ctx.Get(entity.ContextKeyUserID)
	id := userID.(string)
	fmt.Println(code)
	fmt.Println(reportId)
	fmt.Println(id)
	return c.ReportUseCase.GetReport(ctx, code, reportId, id)
}

// CreateReportRequest - 日報作成リクエスト
type CreateReportRequest struct {
	// 本文
	Body string `json:"body"`
	// 実施したタスクリスト
	Tasks []entity.Task `json:"tasks"`
}

// CreateReport godoc
// @Summary  日報作成API
// @Tags     Report
// @Accept   json
// @Produce  json
// @Param    organizationCode path     string               true "組織コード"
// @Param    request          body     CreateReportRequest  true "日報作成リクエスト"
// @Success  201              {object} CreateReportRequest  "Created"
// @Failure  400              {object} entity.ErrorResponse "BadRequest"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Router   /organizations/{organizationCode}/reports [post]
// @Security Bearer
func (c *ReportController) CreateReport(ctx *gin.Context) (interface{}, error) {
	var req CreateReportRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	code := ctx.Params.ByName("organizationCode")
	userID, _ := ctx.Get(entity.ContextKeyUserID)
	id := userID.(string)
	fmt.Println(code)
	fmt.Println(req.Body)
	fmt.Println(id)
	return c.ReportUseCase.CreateReport(ctx, code, id, req.Body, req.Tasks)
}
