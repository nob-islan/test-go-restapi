package model

// サンプルのリクエストモデルです。
type SampleRequest struct {
	Name string `json:"name"` // 名前
}

// サンプルのレスポンスモデルです。
type SampleResponse struct {
	Message string `json:"message"` // メッセージ
}
