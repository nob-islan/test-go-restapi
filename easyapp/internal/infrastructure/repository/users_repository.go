package repository

import (
	"github.com/nob-islan/test-go-restapi/internal/domain"
)

type usersRepository struct {
	//db *sql.DB
}

func NewUsersRepository() domain.UserRepository {
	return &usersRepository{}
}

// 名前を検索条件にユーザ情報をデータベースから抽出します。
func (r *usersRepository) SelectByName(name string) []domain.Users {

	// const sql string = "SELECT * FROM users WHERE name LIKE CONCAT('%', ?, '%')"

	// rows, _ := r.db.Query(sql, name)
	// defer rows.Close()

	// // 返却値作成
	// var users []domain.Users
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	var age int
	// 	rows.Scan(&id, &name, &age)
	// 	users = append(users, domain.NewUsers(id, name, age))
	// }

	return []domain.Users{domain.NewUsers(1, "nob01", 13), domain.NewUsers(2, "nob02", 706)}
}
