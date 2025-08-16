package router

import (
	"net/http"

	"github.com/nob-islan/test-go-restapi/internal/handler"
	"github.com/nob-islan/test-go-restapi/internal/infrastructure/repository"
	"github.com/nob-islan/test-go-restapi/internal/usecase"
)

type usersRouter struct {
	//db *sql.DB
}

func NewUsersRouter() Router {
	return &usersRouter{}
}

func (r *usersRouter) SetRouting(m *http.ServeMux) *http.ServeMux {

	handler := handler.NewUsersHandler(usecase.NewUsersUsecase(repository.NewUsersRepository()))

	m.HandleFunc(basePath+"/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.Search(w, r)
		}
	})

	return m
}
