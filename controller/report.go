package controller

import (
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	Reports []*ReportResponse `json:"reports"`
}

type ReportResponse struct {
	// 日報レスポンス
	ID string `json:"id"`
	// ユーザID
	UserID string `json:"userId"`
	// ユーザ名
	UserName string `json:"userName"`
	// 本文
	Body string `json:"body"`
	// レビュー本文
	ReviewBody *string `json:"reviewBody"`
	// 実施したタスクリスト
	Tasks []entity.Task `json:"tasks"`
	// 作成日時
	Timestamp time.Time `json:"timestamp"`
}

// GetReports godoc
// @Summary  日報リスト取得API
// @Tags    Report
// @Accept  json
// @Produce json
// @Param    organizationCode path     string               true "組織コード"
// @Success  200              {object} ReportsResponse      "OK"
// @Failure  400              {object} entity.ErrorResponse "BadRequest"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure 404      {object} entity.ErrorResponse "Not Found"
// @Router   /organizations/{organizationCode}/reports [get]
// @Security Bearer
func (c *ReportController) GetReports(ctx *gin.Context) (interface{}, error) {
	user, _ := ctx.Get(entity.ContextKeyUser)
	oUser := user.(*entity.OrganizationUser)
	reports, err := c.ReportUseCase.GetReports(ctx, oUser.OrganizationID)
	if err != nil {
		return nil, err
	}
	var reportResponses []*ReportResponse
	for _, report := range reports {
		reportResponses = append(reportResponses, &ReportResponse{
			ID:         report.ID,
			UserID:     report.UserID,
			UserName:   oUser.UserName,
			Body:       report.Body,
			ReviewBody: report.ReviewBody,
			Tasks:      report.Tasks,
			Timestamp:  report.Timestamp,
		})
	}
	return ReportsResponse{Reports: reportResponses}, nil
}

// GetReport godoc
// @Summary  日報取得API
// @Tags     Report
// @Accept   json
// @Produce  json
// @Param    organizationCode path     string               true "組織コード"
// @Param   reportId path string                     true "日報ID"
// @Success  200              {object} ReportResponse       "OK"
// @Failure  401              {object} entity.ErrorResponse "Unauthorized"
// @Failure  403              {object} entity.ErrorResponse "Forbidden"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Router   /organizations/{organizationCode}/reports/{reportId} [get]
// @Security Bearer
func (c *ReportController) GetReport(ctx *gin.Context) (interface{}, error) {
	reportId := ctx.Params.ByName("reportId")
	user, _ := ctx.Get(entity.ContextKeyUser)
	oUser := user.(*entity.OrganizationUser)
	report, err := c.ReportUseCase.GetReport(ctx, oUser.OrganizationID, reportId)
	if err != nil {
		return nil, err
	}
	return &ReportResponse{
		ID:         reportId,
		UserID:     report.UserID,
		UserName:   oUser.UserName,
		Body:       report.Body,
		ReviewBody: report.ReviewBody,
		Tasks:      report.Tasks,
		Timestamp:  report.Timestamp,
	}, nil
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
	user, _ := ctx.Get(entity.ContextKeyUser)
	oUser := user.(*entity.OrganizationUser)
	report, err := c.ReportUseCase.CreateReport(ctx, oUser.OrganizationID, oUser.UserID, req.Body, req.Tasks)
	if err != nil {
		return nil, err
	}
	return CreateReportRequest{
		Body:  report.Body,
		Tasks: report.Tasks,
	}, nil
}

// ReviewReport godoc
// @Summary バッチ処理用の日報レビューAPI
// @Tags     Report
// @Accept   json
// @Produce  json
// @Param    reportId         path     string               true "日報ID"
// @Param   request  body entity.ReviewReportRequest true "日報レビューリクエスト"
// @Success 204      "No Content"
// @Failure  404              {object} entity.ErrorResponse "Not Found"
// @Router  /reports/{reportId} [put]
func (c *ReportController) ReviewReport(ctx *gin.Context) (interface{}, error) {
	var req *entity.ReviewReportRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, entity.NewError(http.StatusBadRequest, err)
	}
	reportID := ctx.Params.ByName("reportId")
	return nil, c.ReportUseCase.ReviewReport(ctx, reportID, req.ReviewBody)
}
