package usecase

import (
	"testing"

	"github.com/nob-islan/test-go-restapi/internal/domain"
	"github.com/nob-islan/test-go-restapi/internal/types"
	"github.com/nob-islan/test-go-restapi/internal/usecase/payload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// モックの定義
type MockUsersRepository struct {
	mock.Mock
}

func (m *MockUsersRepository) SelectByName(username string) []domain.Users {
	args := m.Called(username)
	return args.Get(0).([]domain.Users)
}

func Test_UsersUsecase_Search(t *testing.T) {

	tests := []struct {
		name         string                          // テストケース名
		requestBody  payload.SearchUsersIn           // リクエストボディ
		setupMock    func(mock *MockUsersRepository) // モック設定
		expectedBody payload.SearchUsersOut          // 期待されるレスポンスボディ
	}{
		{
			name:        "success",
			requestBody: payload.NewSearchUsersIn("nob"),
			setupMock: func(mock *MockUsersRepository) {
				mock.On(
					"SelectByName",
					"nob",
				).Return([]domain.Users{
					domain.NewUsers(1, "nob", 13),
				})
			},
			expectedBody: payload.NewSearchUsersOut(
				[]types.UsersDto{
					types.NewUsersDto(1, "nob", 13),
				},
			),
		},
	}

	for _, testcase := range tests {

		// モックリポジトリ初期化
		mockRepository := new(MockUsersRepository)

		t.Run(testcase.name, func(t *testing.T) {
			// モックの期待される動作を定義
			testcase.setupMock(mockRepository)

			// usecaseの初期化
			s := NewUsersUsecase(mockRepository)

			// usecaseの実行
			result := s.Search(testcase.requestBody)

			// レスポンスの検証
			assert.Equal(t, testcase.expectedBody, result)
		})
	}
}
