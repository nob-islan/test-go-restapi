package types

// ユーザ情報格納dtoです。
type UsersDto struct {
	Id   int    `json:"id"`   // 管理番号
	Name string `json:"name"` // 名前
	Age  int    `json:"age"`  // 年齢
}

// ユーザ情報格納dtoのコンストラクタです。
func NewUsersDto(id int, name string, age int) UsersDto {
	return UsersDto{Id: id, Name: name, Age: age}
}
