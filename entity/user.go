package entity

const ContextKeyUser = "user"

type OrganizationUser struct {
	UserID         string `json:"userId"`
	UserName       string `json:"userName"`
	OrganizationID string `json:"organizationId"`
	IsAdmin        bool   `json:"isAdmin"`
}

type User struct {
	// ユーザID
	ID string `json:"id"`
	// ユーザ名
	Name string `json:"name"`
	// メールアドレス
	Email string `json:"email"`
	// 所属する組織リスト
	Organizations []UserOrganization `json:"organizations"`
}
