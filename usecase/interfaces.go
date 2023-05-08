package usecase

import (
	"github.com/fy23-gw-gackathon/reportify-backend/entity"
	"golang.org/x/net/context"
)

type UserRepo interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	GetUsers(ctx context.Context, organizationID string) ([]*entity.User, error)
	GetOrganizationUser(ctx context.Context, organizationCode string, userID string) (*entity.OrganizationUser, error)
	GetUserIDFromToken(ctx context.Context, token string) (*string, error)
	CreateUser(ctx context.Context, email, organizationID string) (*entity.User, error)
	UpdateUserRole(ctx context.Context, organizationID, userID string, role bool) error
	DeleteUser(ctx context.Context, organizationID, userID string) error
}

type OrganizationRepo interface {
	GetOrganization(ctx context.Context, organizationID string) (*entity.Organization, error)
	GetOrganizations(ctx context.Context, userID string) ([]*entity.Organization, error)
	UpdateOrganization(ctx context.Context, organizationID, organizationName, organizationCode, mission, vision, value string) (*entity.Organization, error)
}

type ReportRepo interface {
	GetReport(ctx context.Context, organizationID *string, reportID string) (*entity.Report, error)
	GetReports(ctx context.Context, organizationID string) ([]*entity.Report, error)
	CreateReport(ctx context.Context, organizationID, userID string, body string, task []entity.Task) (*entity.Report, error)
	UpdateReviewBody(ctx context.Context, reportID string, reviewBody string) error
	DispatchReport(ctx context.Context, reportID, body string, mvv entity.Mvv) error
}
