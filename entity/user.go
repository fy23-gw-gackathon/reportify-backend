package entity

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
