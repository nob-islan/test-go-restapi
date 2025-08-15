package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nob-islan/test-go-restapi/internal/handler/model"
	"github.com/nob-islan/test-go-restapi/internal/usecase"
	"github.com/nob-islan/test-go-restapi/internal/usecase/payload"
)

// ユーザ情報handlerのインターフェースです。
type UsersHandler interface {
	Search(w http.ResponseWriter, r *http.Request)
}

type usersHandler struct {
	usersUsecase usecase.UsersUsecase
}

// ユーザ情報handlerのコンストラクタです。
func NewUsersHandler(u usecase.UsersUsecase) UsersHandler {
	return &usersHandler{usersUsecase: u}
}

// ユーザを検索します。
func (h *usersHandler) Search(w http.ResponseWriter, r *http.Request) {

	req := model.NewSearchUsersReq(r)
	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out := h.usersUsecase.Search(payload.NewSearchUsersIn(req.Name))

	res := model.NewSearchUsersRes(out.UsersDtos())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
