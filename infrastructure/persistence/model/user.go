package model

import (
	"reportify-backend/entity"
	"time"
)

type User struct {
	ID                string
	Name              string
	Email             string
	CognitoID         string
	Reports           []*Report
	OrganizationUsers []*OrganizationUser
	Organizations     []*Organization `gorm:"many2many:r_organization_users;"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (m User) ToEntity() *entity.User {
	var organizations []entity.UserOrganization
	for _, ou := range m.OrganizationUsers {
		organizations = append(organizations, *ou.ToEntity())
	}
	return &entity.User{
		ID:            m.ID,
		Name:          m.Name,
		Email:         m.Email,
		Organizations: organizations,
	}
}
