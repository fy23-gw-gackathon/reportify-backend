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

func (u ReportUseCase) GetReports(ctx context.Context, organizationCode, userID string) ([]*entity.Report, error) {
	//_, err := u.UserRepo.GetOrganizationUserRole(ctx, organizationCode, &userID, nil)
	//if err != nil {
	//	return nil, entity.NewError(http.StatusUnauthorized, err)
	//}
	return u.ReportRepo.GetReports(ctx, organizationCode, userID)
}

func (u ReportUseCase) GetReport(ctx context.Context, organizationCode, reportID, userID string) (*entity.Report, error) {
	return u.ReportRepo.GetReport(ctx, organizationCode, reportID)
}

func (u ReportUseCase) CreateReport(ctx context.Context, organizationCode, userID, body string, task []entity.Task) (*entity.Report, error) {
	orgID := "01GZR2TYVGFJKWH35BF2J5Z38E"
	return u.ReportRepo.CreateReport(ctx, orgID, userID, body, task)
}
