package router

import (
	"net/http"

	"github.com/nob-islan/test-go-restapi/internal/handler"
)

type sampleRouter struct{}

func NewSampleRouter() Router {
	return &sampleRouter{}
}

func (r *sampleRouter) SetRouting(m *http.ServeMux) *http.ServeMux {

	h := handler.NewSampleHandler()

	// カスタムルータ
	m.HandleFunc(basePath+"/sample", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.Greet(w, r)
		}
	})

	return m
}
