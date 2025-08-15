package util

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

// テスト用データベースに接続します。
func ConnectTestDB(t *testing.T, uri string) *sql.DB {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	// schema.sql を読み込み・実行
	schema, err := os.ReadFile(fmt.Sprintf("test/data/%s/schema.sql", uri))
	if err != nil {
		t.Fatalf("failed to read schema: %v", err)
	}
	_, err = db.Exec(string(schema))
	if err != nil {
		t.Fatalf("failed to execute schema: %v", err)
	}

	// data.sql を読み込み・実行
	data, err := os.ReadFile(fmt.Sprintf("test/data/%s/data.sql", uri))
	if err != nil {
		t.Fatalf("failed to read data: %v", err)
	}
	_, err = db.Exec(string(data))
	if err != nil {
		t.Fatalf("failed to execute data: %v", err)
	}

	return db
}
