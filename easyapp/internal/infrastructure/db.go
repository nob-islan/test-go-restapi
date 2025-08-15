package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// データベースに接続します。
func ConnectDB() *sql.DB {

	const (
		user       string = "root"
		password   string = "password"
		domain     string = "localhost:3306"
		dbName     string = "snaildb"
		driverName string = "mysql"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, domain, dbName)
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatal("Fail to connect to Database")
	}

	// 実際に接続できるかを確認
	err = db.Ping()
	if err != nil {
		log.Fatal("Fail to connect to Database")
	}

	return db
}
