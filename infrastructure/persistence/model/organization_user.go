package model

import "reportify-backend/entity"

type OrganizationUser struct {
	UserID         string
	OrganizationID string
	Role           uint8
	User           *User
	Organization   *Organization
}

func (m OrganizationUser) TableName() string {
	return "r_organization_users"
}

func (m OrganizationUser) ToEntity() *entity.UserOrganization {
	return &entity.UserOrganization{
		ID:      m.OrganizationID,
		IsAdmin: m.Role == 1,
	}
}
