package controller

import (
	"golang.org/x/net/context"
	"reportify-backend/entity"
)

type UserUseCase interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	GetUsers(ctx context.Context, organizationCode string) ([]*entity.User, error)
	InviteUser(ctx context.Context, email, organizationCode, userID string) (*entity.User, error)
}

type OrganizationUseCase interface {
	GetOrganization(ctx context.Context, organizationCode, userID string) (*entity.Organization, error)
	GetOrganizations(ctx context.Context, userID string) ([]*entity.Organization, error)
	UpdateOrganization(ctx context.Context, oldOrganizationCode, userID, organizationName, organizationCode, mission, vision, value string) (*entity.Organization, error)
}

type ReportUseCase interface {
	GetReports(ctx context.Context, organizationCode, userID string) ([]*entity.Report, error)
	GetReport(ctx context.Context, organizationCode, reportID, userID string) (*entity.Report, error)
	CreateReport(ctx context.Context, organizationCode, userID, body string, task []entity.Task) (*entity.Report, error)
}
