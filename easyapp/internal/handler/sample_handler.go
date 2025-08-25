package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nob-islan/test-go-restapi/internal/handler/model"
)

// サンプルハンドラのインターフェースです。
type SampleHandler interface {
	Greet(w http.ResponseWriter, r *http.Request)
}

// サンプルハンドラの構造体です。
type sampleHandler struct{}

// サンプルハンドラのコンストラクタです。
func NewSampleHandler() SampleHandler {
	return &sampleHandler{}
}

// サンプルの挨拶関数です。
func (h *sampleHandler) Greet(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	var res model.SampleResponse

	if name == "" {
		res = model.SampleResponse{Message: "Hello, John doe!"}
	} else {
		res = model.SampleResponse{Message: "Hello, " + name + "!"}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
