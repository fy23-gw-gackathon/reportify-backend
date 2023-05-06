package model

type OrganizationUser struct {
	UserID         string
	OrganizationID string
	role           uint8
	User           *User
	Organization   *Organization
}

func (m OrganizationUser) TableName() string {
	return "r_organization_users"
}
