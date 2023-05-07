package usecase

import (
	"golang.org/x/net/context"
	"reportify-backend/entity"
)

type UserRepo interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	GetUsers(ctx context.Context, organizationID string) ([]*entity.User, error)
	GetOrganizationUser(ctx context.Context, organizationCode string, userID string) (*entity.OrganizationUser, error)
	GetUserIDFromToken(ctx context.Context, token string) (*string, error)
	CreateUser(ctx context.Context, email, organizationID string) (*entity.User, error)
}

type OrganizationRepo interface {
	GetOrganization(ctx context.Context, organizationID string) (*entity.Organization, error)
	GetOrganizations(ctx context.Context, userID string) ([]*entity.Organization, error)
	UpdateOrganization(ctx context.Context, organizationID, organizationName, organizationCode, mission, vision, value string) (*entity.Organization, error)
}

type ReportRepo interface {
	GetReport(ctx context.Context, organizationID, reportID string) (*entity.Report, error)
	GetReports(ctx context.Context, organizationID string) ([]*entity.Report, error)
	CreateReport(ctx context.Context, organizationID, userID string, body string, task []entity.Task) (*entity.Report, error)
}
