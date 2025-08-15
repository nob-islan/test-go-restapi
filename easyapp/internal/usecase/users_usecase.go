package usecase

import (
	"github.com/nob-islan/test-go-restapi/internal/domain"
	"github.com/nob-islan/test-go-restapi/internal/types"
	"github.com/nob-islan/test-go-restapi/internal/usecase/payload"
)

// ユーザ情報usecaseのインターフェースです。
type UsersUsecase interface {
	Search(in payload.SearchUsersIn) payload.SearchUsersOut
}

type usersUsecase struct {
	usersRepository domain.UserRepository
}

// ユーザ情報usecaseのコンストラクタです。
func NewUsersUsecase(r domain.UserRepository) UsersUsecase {
	return &usersUsecase{usersRepository: r}
}

// ユーザを検索します。
func (u *usersUsecase) Search(in payload.SearchUsersIn) payload.SearchUsersOut {

	var usersDtos []types.UsersDto
	for _, user := range u.usersRepository.SelectByName(in.Name()) {
		usersDtos = append(usersDtos, types.NewUsersDto(user.Id(), user.Name(), user.Age()))
	}

	return payload.NewSearchUsersOut(usersDtos)
}
