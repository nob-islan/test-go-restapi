package domain

// ユーザ情報を表すdomain構造体です。
type Users struct {
	id   int    // 管理番号
	name string // 名前
	age  int    // 年齢
}

// ユーザ情報を表すdomain構造体のコンストラクタです。
func NewUsers(id int, name string, age int) Users {
	return Users{id: id, name: name, age: age}
}

func (u *Users) Id() int {
	return u.id
}

func (u *Users) Name() string {
	return u.name
}

func (u *Users) Age() int {
	return u.age
}

// ユーザ情報を取得するrepositoryのインターフェースです。
type UserRepository interface {
	SelectByName(name string) []Users
}
