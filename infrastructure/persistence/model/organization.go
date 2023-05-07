package model

import (
	"reportify-backend/entity"
	"time"
)

type Organization struct {
	ID                string
	Name              string
	Code              string
	Mission           string
	Vision            string
	Value             string
	OrganizationUsers []*OrganizationUser
	Reports           []*Report
	Users             []*User `gorm:"many2many:r_organization_users;"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (m Organization) ToEntity() *entity.Organization {
	return &entity.Organization{
		ID:   m.ID,
		Name: m.Name,
		Code: m.Code,
		Mvv: entity.Mvv{
			Mission: m.Mission,
			Vision:  m.Vision,
			Value:   m.Value,
		},
	}
}
