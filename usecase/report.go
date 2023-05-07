package usecase

import (
	"golang.org/x/net/context"
	"reportify-backend/entity"
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
	return u.ReportRepo.GetReport(ctx, organizationID, reportID)
}

func (u ReportUseCase) CreateReport(ctx context.Context, organizationID, userID, body string, task []entity.Task) (*entity.Report, error) {
	return u.ReportRepo.CreateReport(ctx, organizationID, userID, body, task)
}
