package usecase

import (
	"golang.org/x/net/context"
	"reportify-backend/entity"
)

type UserRepo interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	GetUsers(ctx context.Context, organizationCode string) ([]*entity.User, error)
	GetUserIDFromToken(ctx context.Context, token string) (*string, error)
	CreateUser(ctx context.Context, email, organizationID string) (*entity.User, error)
	GetOrganizationUserRole(ctx context.Context, organizationCode string, userID, email *string) (*entity.OrganizationUser, error)
}

type OrganizationRepo interface {
	GetOrganizations(ctx context.Context, userID string) ([]*entity.Organization, error)
	GetOrganization(ctx context.Context, organizationCode, userID string) (*entity.Organization, error)
	UpdateOrganization(ctx context.Context, oldOrganizationCode, organizationName, organizationCode, mission, vision, value string) (*entity.Organization, error)
}

type ReportRepo interface {
	GetReports(ctx context.Context, organizationCode, userID string) ([]*entity.Report, error)
	GetReport(ctx context.Context, organizationCode, reportID string) (*entity.Report, error)
	CreateReport(ctx context.Context, organizationCode, userID, body string, task []entity.Task) (*entity.Report, error)
}
