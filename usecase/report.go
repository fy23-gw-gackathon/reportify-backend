package usecase

import (
	"errors"
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"golang.org/x/net/context"
	"net/http"
)

type ReportUseCase struct {
	ReportRepo
	UserRepo
}

func NewReportUseCase(reportRepo ReportRepo, userRepo UserRepo) *ReportUseCase {
	return &ReportUseCase{reportRepo, userRepo}
}

func (u ReportUseCase) GetReports(ctx context.Context, organizationID string) ([]*entity.Report, error) {
	return u.ReportRepo.GetReports(ctx, organizationID)
}

func (u ReportUseCase) GetReport(ctx context.Context, organizationID, reportID string) (*entity.Report, error) {
	return u.ReportRepo.GetReport(ctx, &organizationID, reportID)
}

func (u ReportUseCase) CreateReport(ctx context.Context, organizationID, userID, body string, task []entity.Task) (*entity.Report, error) {
	reports, err := u.ReportRepo.CreateReport(ctx, organizationID, userID, body, task)
	if err != nil {
		return nil, err
	}
	return reports, u.ReportRepo.DispatchReport(ctx, reports.ID, body)
}

func (u ReportUseCase) ReviewReport(ctx context.Context, reportID, reviewBody string) error {
	report, err := u.ReportRepo.GetReport(ctx, nil, reportID)
	if err != nil {
		return err
	}

	if report.ReviewBody != nil {
		return entity.NewError(http.StatusConflict, errors.New("already reviewed"))
	}

	return u.ReportRepo.UpdateReviewBody(ctx, reportID, reviewBody)
}
