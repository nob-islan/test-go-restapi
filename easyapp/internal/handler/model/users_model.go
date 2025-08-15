package model

import (
	"net/http"

	"github.com/nob-islan/test-go-restapi/internal/types"
)

// ユーザ検索向けリクエストモデルです。
type SearchUsersReq struct {
	Name string `json:"name" example:"nob"` // ユーザ名
}

// ユーザ検索向けリクエストモデルのコンストラクタです。
func NewSearchUsersReq(r *http.Request) SearchUsersReq {
	return SearchUsersReq{Name: r.URL.Query().Get("name")}
}

// ユーザ検索向けレスポンスモデルです。
type SearchUsersRes struct {
	Users []types.UsersDto `json:"users"` // ユーザ情報
}

// ユーザ検索向けレスポンスモデルのコンストラクタです。
func NewSearchUsersRes(users []types.UsersDto) SearchUsersRes {
	return SearchUsersRes{Users: users}
}
