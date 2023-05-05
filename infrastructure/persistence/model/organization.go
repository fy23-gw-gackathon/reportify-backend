package model

import "time"

type Organization struct {
	ID                string
	Name              string
	Code              string
	Mission           string
	Vision            string
	Value             string
	OrganizationUsers []*OrganizationUser
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
