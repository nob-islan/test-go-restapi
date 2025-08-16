package router

import (
	"net/http"
)

// routerのインターフェースです。
type Router interface {
	SetRouting(m *http.ServeMux) *http.ServeMux
}

// APIのベースURI
const basePath string = "/api/v1"

// ルーティングを設定します。
func Routing() *http.ServeMux {

	// // データベースに接続
	// db := infrastructure.ConnectDB()

	// 各handlerに紐づくルーティングを設定
	m := http.NewServeMux()

	// ユーザ情報
	m = NewUsersRouter().SetRouting(m)

	return m
}
