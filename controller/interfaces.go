package controller

import (
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"golang.org/x/net/context"
)

type UserUseCase interface {
	GetUserFromToken(ctx context.Context, token string) (*entity.User, error)
	GetUsers(ctx context.Context, organizationID string) ([]*entity.User, error)
	InviteUser(ctx context.Context, email, organizationID string) (*entity.User, error)
}

type OrganizationUseCase interface {
	GetOrganization(ctx context.Context, organizationID string) (*entity.Organization, error)
	GetOrganizations(ctx context.Context, userID string) ([]*entity.Organization, error)
	UpdateOrganization(ctx context.Context, organizationID, organizationName, organizationCode, mission, vision, value string) (*entity.Organization, error)
}

type ReportUseCase interface {
	GetReport(ctx context.Context, organizationID, reportID string) (*entity.Report, error)
	GetReports(ctx context.Context, organizationID string) ([]*entity.Report, error)
	CreateReport(ctx context.Context, organizationID, userID, body string, task []entity.Task) (*entity.Report, error)
	ReviewReport(ctx context.Context, reportId, reviewBody string) error
}
