package router

import "net/http"

// routerのインターフェースです。
type Router interface {
	SetRouting(m *http.ServeMux) *http.ServeMux
}

// APIのベースURI
const basePath string = "/api/v1"

// ルーティングを設定します。
func Routing() *http.ServeMux {

	// 各handlerに紐づくルーティングを設定
	m := http.NewServeMux()

	// sample
	m = NewSampleRouter().SetRouting(m)

	return m
}
