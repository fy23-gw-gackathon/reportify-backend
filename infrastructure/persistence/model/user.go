package model

import (
	"time"
)

type User struct {
	ID                string
	Name              string
	Email             string
	CognitoID         string
	Reports           []*Report
	OrganizationUsers []*OrganizationUser
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
