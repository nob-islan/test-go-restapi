package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nob-islan/test-go-restapi/internal/types"
	"github.com/nob-islan/test-go-restapi/internal/usecase/payload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// モックの定義
type MockUsersUsecase struct {
	mock.Mock
}

func (m *MockUsersUsecase) Search(in payload.SearchUsersIn) payload.SearchUsersOut {
	args := m.Called(in)
	return args.Get(0).(payload.SearchUsersOut)
}

func Test_UsersHandler_Search(t *testing.T) {

	tests := []struct {
		name               string                       // テストケース名
		requestParam       map[string]string            // リクエストパラメータ
		setupMock          func(mock *MockUsersUsecase) // モック設定
		expectedStatusCode int                          // 期待されるHTTPステータス
		expectedBody       string                       // 期待されるレスポンスボディ
	}{
		{
			name:         "success",
			requestParam: map[string]string{"name": "nob"},
			setupMock: func(mock *MockUsersUsecase) {
				mock.On(
					"Search",
					payload.NewSearchUsersIn("nob"),
				).Return(payload.NewSearchUsersOut([]types.UsersDto{
					{
						Id:   1,
						Name: "nob",
						Age:  13,
					},
				}))
			},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"users":[{"id":1,"name":"nob","age":13}]}`,
		},
		{
			name:               "invalid request",
			requestParam:       map[string]string{"name": ""},
			setupMock:          func(mock *MockUsersUsecase) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "",
		},
	}

	for _, testcase := range tests {

		// モックusecase初期化
		mockUsecase := new(MockUsersUsecase)

		t.Run(testcase.name, func(t *testing.T) {
			// モックの期待される動作を定義
			testcase.setupMock(mockUsecase)

			// handlerの初期化
			h := NewUsersHandler(mockUsecase)

			// リクエストとレスポンスの準備
			uri := "/users?name=" + testcase.requestParam["name"]
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			res := httptest.NewRecorder()

			// handlerの実行
			h.Search(res, req)

			// レスポンスの検証
			assert.Equal(t, testcase.expectedStatusCode, res.Code)
			if testcase.expectedBody != "" {
				assert.JSONEq(t, testcase.expectedBody, res.Body.String())
			}
		})
	}
}
