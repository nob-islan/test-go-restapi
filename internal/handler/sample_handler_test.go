package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SampleHandler_Greetのテスト
func Test_SampleHandler_Greet(t *testing.T) {

	tests := []struct {
		name               string            // テストケース名
		requestParam       map[string]string // リクエストパラメータ
		expectedStatusCode int               // 期待されるHTTPステータス
		expectedBody       string            // 期待されるレスポンスボディ
	}{
		{
			name:               "success",
			requestParam:       map[string]string{"name": "testnob"},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"message":"Hello, testnob!"}`,
		},
		{
			name:               "usecase error",
			requestParam:       map[string]string{"name": ""},
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"message":"Hello, John doe!"}`,
		},
	}

	for _, testcase := range tests {

		t.Run(testcase.name, func(t *testing.T) {
			// handlerの初期化
			h := NewSampleHandler()

			// リクエストとレスポンスの準備
			uri := "/sample?name=" + testcase.requestParam["name"]
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			res := httptest.NewRecorder()

			// handlerの実行
			h.Greet(res, req)

			// レスポンスの検証
			assert.Equal(t, testcase.expectedStatusCode, res.Code)
			if testcase.expectedBody != "" {
				assert.JSONEq(t, testcase.expectedBody, res.Body.String())
			}
		})
	}
}
