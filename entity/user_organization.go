package entity

// UserOrganization - ユーザが所属する組織
type UserOrganization struct {
	// 組織ID
	ID string `json:"id"`
	// ロール
	IsAdmin bool `json:"is_admin"`
}
