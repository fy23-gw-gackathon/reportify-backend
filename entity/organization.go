package entity

type Organization struct {
	// 組織ID
	Id string `json:"id"`
	// 組織名
	Name string `json:"name"`
	// 組織コード
	Code string `json:"code"`
	Mvv  Mvv    `json:"mvv"`
}

type Mvv struct {
	// ミッション
	Mission string `json:"mission"`
	// ビジョン
	Vision string `json:"vision"`
	// バリュー
	Value string `json:"value"`
}
