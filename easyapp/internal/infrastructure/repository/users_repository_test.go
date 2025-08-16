package repository

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nob-islan/test-go-restapi/internal/domain"
	"github.com/stretchr/testify/assert"
)

// UserinfoRepository_SelectByUsernameのテスト
func Test_UserinfoRepository_SelectByName(t *testing.T) {

	tests := []struct {
		name         string           // テストケース名
		queryParam   string           // クエリパラメータ
		setup        func(db *sql.DB) // 事前セットアップ関数
		expectedBody []domain.Users   // 期待されるレスポンスボディ
	}{
		{
			name:       "success",
			queryParam: "nob",
			setup:      func(db *sql.DB) {},
			expectedBody: []domain.Users{
				domain.NewUsers(1, "nob01", 13),
				domain.NewUsers(2, "nob02", 706),
			},
		},
	}

	for _, testcase := range tests {

		t.Run(testcase.name, func(t *testing.T) {

			// テストデータベースおよびrepository初期化
			// db := util.ConnectTestDB(t, "users")
			r := NewUsersRepository()

			// // 事前セットアップ
			// testcase.setup(db)

			// repositoryの実行
			result := r.SelectByName(testcase.queryParam)

			// レスポンスの確認
			assert.Equal(t, testcase.expectedBody, result)
		})
	}
}
