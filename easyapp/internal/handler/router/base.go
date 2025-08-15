package router

import (
	"net/http"

	"github.com/nob-islan/test-go-restapi/internal/infrastructure"
)

// routerのインターフェースです。
type Router interface {
	SetRouting(m *http.ServeMux) *http.ServeMux
}

// APIのベースURI
const basePath string = "/api/v1"

// ルーティングを設定します。
func Routing() *http.ServeMux {

	// データベースに接続
	db := infrastructure.ConnectDB()

	// 各handlerに紐づくルーティングを設定
	m := http.NewServeMux()

	// ユーザ情報
	m = NewUsersRouter(db).SetRouting(m)

	return m
}
