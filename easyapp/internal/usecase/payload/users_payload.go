package payload

import "github.com/nob-islan/test-go-restapi/internal/types"

// ユーザ検索処理向け入力モデルです。
type SearchUsersIn struct {
	name string // 名前
}

// ユーザ向け検索処理向け入力モデルのコンストラクタです。
func NewSearchUsersIn(name string) SearchUsersIn {
	return SearchUsersIn{name: name}
}

func (i *SearchUsersIn) Name() string {
	return i.name
}

// ユーザ検索処理向け出力モデルです。
type SearchUsersOut struct {
	usersDtos []types.UsersDto // ユーザ情報格納dtoのスライス
}

// ユーザ検索処理向け出力モデルのコンストラクタです。
func NewSearchUsersOut(usersDtos []types.UsersDto) SearchUsersOut {
	return SearchUsersOut{usersDtos: usersDtos}
}

func (o *SearchUsersOut) UsersDtos() []types.UsersDto {
	return o.usersDtos
}
